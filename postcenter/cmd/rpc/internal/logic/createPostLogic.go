package logic

import (
	"context"
	"errors"

	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"
	"go-zero_less/postcenter/model"
	"go-zero_less/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *pb.PostInfoCreateRequest) (*pb.PostInfoId, error) {
	user, err := l.svcCtx.UserCenterClient.GetUser(l.ctx, &usercenter.UserInfoOneReq{
		Id: in.PostInfo.UserId,
	})
	if user == nil || errors.Is(err, sqlx.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "文章创建失败,用户不存在：%v", err)
	}

	post := model.Posts{
		Title:   in.PostInfo.Title,
		Content: in.PostInfo.Content,
		UserId:  in.PostInfo.UserId,
	}
	res, err := l.svcCtx.PostModel.Insert(l.ctx, &post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "文章创建失败：%v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "文章创建失败：%v", err)
	}

	//更新用户文章数
	_, err = l.svcCtx.UserCenterClient.UserPostNumIncr(l.ctx, &usercenter.UserPostNumReq{
		Id:      in.PostInfo.UserId,
		PostNum: 1,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "<UNK>%v", err)
	}

	return &pb.PostInfoId{
		Id: uint64(id),
	}, nil
}
