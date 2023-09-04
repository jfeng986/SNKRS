package logic

import (
	"context"

	"snkrs/apps/api/internal/svc"
	"snkrs/apps/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LimitedSaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLimitedSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LimitedSaleLogic {
	return &LimitedSaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LimitedSaleLogic) LimitedSale() (resp *types.LimitedSaleResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
