package logic

import (
	"context"
	"fmt"

	"snkrs/apps/rpc_service/product/internal/svc"
	"snkrs/apps/rpc_service/product/product"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckAndUpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	luaCheckAndUpdateScript = `
		local counts = redis.call("HMGET", KEYS[1], "total", "reserve")
		local total = tonumber(counts[1])
		local reserve = tonumber(counts[2])
		if reserve + 1 <= total then
			redis.call("HINCRBY", KEYS[1], "reserve", 1)
			return 1
		end
		return 0
		`
)

func NewCheckAndUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndUpdateStockLogic {
	return &CheckAndUpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAndUpdateStockLogic) CheckAndUpdateStock(in *product.CheckAndUpdateStockRequest) (*product.CheckAndUpdateStockResponse, error) {
	val, err := l.svcCtx.BizRedis.EvalCtx(l.ctx, luaCheckAndUpdateScript, []string{stockKey(in.ProductId)})
	if err != nil {
		return nil, err
	}
	if val.(int64) == 0 {
		return nil, status.Errorf(codes.ResourceExhausted, fmt.Sprintf("insufficient stock: %d", in.ProductId))
	}
	return &product.CheckAndUpdateStockResponse{}, nil
}

func stockKey(pid int64) string {
	return fmt.Sprintf("stock:%d", pid)
}
