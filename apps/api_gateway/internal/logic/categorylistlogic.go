package logic

import (
	"context"

	"snkrs/apps/api_gateway/internal/svc"
	"snkrs/apps/api_gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList(req *types.CategoryListRequest) (resp *types.CategoryListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
