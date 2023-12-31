package logic

import (
	"context"
	"fmt"
	"sync"

	"snkrs/apps/rpc_service/product/internal/model"
	"snkrs/apps/rpc_service/product/internal/svc"
	"snkrs/apps/rpc_service/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var productMutex sync.Mutex

func (l *ProductLogic) Product(in *product.ProductItemRequest) (*product.ProductItem, error) {
	productMutex.Lock()
	defer productMutex.Unlock()

	v, err, _ := l.svcCtx.SingleGroup.Do(fmt.Sprintf("product:%d", in.ProductId), func() (interface{}, error) {
		return l.svcCtx.ProductModel.FindOne(l.ctx, in.ProductId)
	})
	if err != nil {
		return nil, err
	}
	p := v.(*model.Product)
	return &product.ProductItem{
		ProductId: p.Id,
		Name:      p.Name,
		Stock:     p.Stock,
	}, nil
}
