package logic

import (
	"context"

	"snkrs/apps/rpc_service/product/internal/svc"
	"snkrs/apps/rpc_service/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationProductsLogic struct {
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	productListLogic *ProductListLogic

	logx.Logger
}

func NewOperationProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationProductsLogic {
	return &OperationProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	validStatus          = 1
	operationProductsKey = "operation#products"
)

func (l *OperationProductsLogic) OperationProducts(in *product.OperationProductsRequest) (*product.OperationProductsResponse, error) {
	opProducts, ok := l.svcCtx.LocalCache.Get(operationProductsKey)
	if ok {
		return &product.OperationProductsResponse{Products: opProducts.([]*product.ProductItem)}, nil
	}

	pos, err := l.svcCtx.OperationModel.OperationProducts(l.ctx, validStatus)
	if err != nil {
		return nil, err
	}
	var pids []int64
	for _, p := range pos {
		pids = append(pids, p.ProductId)
	}
	products, err := l.productListLogic.productsByIds(l.ctx, pids)
	if err != nil {
		return nil, err
	}
	var pItems []*product.ProductItem
	for _, p := range products {
		pItems = append(pItems, &product.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
		})
	}
	l.svcCtx.LocalCache.Set(operationProductsKey, pItems)
	return &product.OperationProductsResponse{Products: pItems}, nil
}
