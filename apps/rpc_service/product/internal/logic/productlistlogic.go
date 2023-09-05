package logic

import (
	"context"

	"snkrs/apps/rpc_service/product/internal/svc"
	"snkrs/apps/rpc_service/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductListLogic) ProductList(in *product.ProductListRequest) (*product.ProductListResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ProductListResponse{}, nil
}
