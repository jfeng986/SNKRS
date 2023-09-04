package logic

import (
	"context"

	"snkrs/apps/api/internal/svc"
	"snkrs/apps/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartListLogic) CartList(req *types.CartListRequest) (resp *types.CartListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}