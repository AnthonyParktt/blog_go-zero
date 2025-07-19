package commentsinfologic

import (
	"context"

	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentsInfoLogic {
	return &DeleteCommentsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentsInfoLogic) DeleteCommentsInfo(in *pb.CommentsInfoId) (*pb.CommentsInfoDeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CommentsInfoDeleteResponse{}, nil
}
