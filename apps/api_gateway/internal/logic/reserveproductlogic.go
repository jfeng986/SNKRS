package logic

import (
	"context"

	"snkrs/apps/api_gateway/internal/svc"
	"snkrs/apps/api_gateway/internal/types"
	"snkrs/apps/rpc_service/reserve/rpc/reserve"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReserveProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReserveProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReserveProductLogic {
	return &ReserveProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReserveProductLogic) ReserveProduct(req *types.ReserveProductRequest) (resp *types.ReserveProductResponse, err error) {
	_, err = l.svcCtx.ReserveRPC.ReserveOrder(l.ctx, &reserve.ReserveOrderRequest{UserId: req.UserID, ProductId: req.ProductID})
	if err != nil {
		return &types.ReserveProductResponse{Message: "Failed"}, err
	}
	return &types.ReserveProductResponse{Message: "Success"}, nil
}
