package logic

import (
	"context"

	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"
	"go-zero_less/postcenter/model"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostLogic) UpdatePost(in *pb.PostInfoBase) (*pb.ExecRows, error) {
	if in.Id == 0 {
		return nil, status.Error(codes.DataLoss, "invalid post id")
	}
	postModel := &model.Posts{
		Id:      in.Id,
		Title:   in.Title,
		Content: in.Content,
		UserId:  in.UserId,
	}
	err := l.svcCtx.PostModel.Update(l.ctx, postModel)
	if err != nil {
		return nil, err
	}

	return &pb.ExecRows{Rows: 1}, nil
}
