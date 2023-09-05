package logic

import (
	"context"

	"snkrs/apps/comment/rpc/comment"
	"snkrs/apps/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentsLogic {
	return &CommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentsLogic) Comments(in *comment.CommentsRequest) (*comment.CommentsResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.CommentsResponse{}, nil
}
