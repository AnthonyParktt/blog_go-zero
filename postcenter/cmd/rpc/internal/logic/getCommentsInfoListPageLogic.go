package logic

import (
	"context"

	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsInfoListPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsInfoListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsInfoListPageLogic {
	return &GetCommentsInfoListPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsInfoListPageLogic) GetCommentsInfoListPage(in *pb.CommentPostId) (*pb.CommentsInfoListPage, error) {
	// todo: add your logic here and delete this line

	return &pb.CommentsInfoListPage{}, nil
}
