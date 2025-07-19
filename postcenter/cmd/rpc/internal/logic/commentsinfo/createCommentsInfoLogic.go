package commentsinfologic

import (
	"context"

	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentsInfoLogic {
	return &CreateCommentsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentsInfoLogic) CreateCommentsInfo(in *pb.CommentsInfoCreateRequest) (*pb.CommentsInfoId, error) {
	// todo: add your logic here and delete this line

	return &pb.CommentsInfoId{}, nil
}
