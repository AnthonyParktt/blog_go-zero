package logic

import (
	"context"

	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsInfoDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsInfoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsInfoDetailLogic {
	return &GetCommentsInfoDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsInfoDetailLogic) GetCommentsInfoDetail(in *pb.CommentsInfoId) (*pb.CommentsInfoDetail, error) {
	// todo: add your logic here and delete this line

	return &pb.CommentsInfoDetail{}, nil
}
