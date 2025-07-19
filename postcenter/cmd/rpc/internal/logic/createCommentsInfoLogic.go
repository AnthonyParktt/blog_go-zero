package logic

import (
	"context"
	"github.com/spf13/cast"
	"go-zero_less/postcenter/model"

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
	res, err := l.svcCtx.CommentsModel.Insert(l.ctx, &model.Comments{
		PostId:  in.CommentsInfo.PostId,
		UserId:  in.CommentsInfo.UserId,
		Content: in.CommentsInfo.Content,
	})
	if err != nil {
		return nil, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &pb.CommentsInfoId{
		Id: cast.ToUint64(lastId),
	}, nil
}
