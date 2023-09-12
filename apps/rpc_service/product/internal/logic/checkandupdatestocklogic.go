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

/*
HMGET fetches the current values of "total" and "reserve" from the hash table identified by KEYS[1].
These values are converted to numbers.
An if condition checks if "reserve + 1" is less than or equal to "total".
If the condition is true, it increments "reserve" by 1 using HINCRBY.
If the increment is successful, the script returns 1.
If the condition is not met, the script returns 0.
*/

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
