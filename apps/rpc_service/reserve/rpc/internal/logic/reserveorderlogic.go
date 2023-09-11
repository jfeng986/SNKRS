package logic

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"snkrs/apps/rpc_service/product/product"
	"snkrs/apps/rpc_service/reserve/rpc/internal/svc"
	"snkrs/apps/rpc_service/reserve/rpc/reserve"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReserveOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	limiter    *limit.PeriodLimit
	localCache *collection.Cache
}

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

const (
	limitPeriod       = 10
	limitQuota        = 1
	reserveUserPrefix = "reserve#u#"
	localCacheExpire  = time.Second * 60
)

func NewReserveOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReserveOrderLogic {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	return &ReserveOrderLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		localCache: localCache,
		limiter:    limit.NewPeriodLimit(limitPeriod, limitQuota, svcCtx.BizRedis, reserveUserPrefix),
	}
}

func (l *ReserveOrderLogic) ReserveOrder(in *reserve.ReserveOrderRequest) (*reserve.ReserveOrderResponse, error) {
	code, _ := l.limiter.Take(strconv.FormatInt(in.UserId, 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "Number of requests exceeded the limit")
	}
	p, err := l.svcCtx.ProductRPC.Product(l.ctx, &product.ProductItemRequest{ProductId: in.ProductId})
	if err != nil {
		return nil, err
	}
	if p.Stock <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient stock")
	}
	kd, err := json.Marshal(&KafkaData{Uid: in.UserId, Pid: in.ProductId})
	if err != nil {
		return nil, err
	}
	if err := l.svcCtx.KafkaPusher.Push(string(kd)); err != nil {
		return nil, err
	}
	return &reserve.ReserveOrderResponse{}, nil
}
