package postinfologic

import (
	"context"
	"errors"
	"go-zero_less/postcenter/model"
	"go-zero_less/usercenter/cmd/rpc/usercenter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostLogic) DeletePost(in *pb.PostInfoId) (*pb.PostInfoDeleteResponse, error) {
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "文章不存在：%v", err)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部错误%v", err)
	}

	_, err = l.svcCtx.PostModel.DeleteSoft(l.ctx, post.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除文章异常%v", err)
	}

	//更新用户文章数
	res, err := l.svcCtx.UserCenterClient.UserPostNumIncr(l.ctx, &usercenter.UserPostNumReq{
		Id:      post.UserId,
		PostNum: -1,
	})
	if err != nil || res.Result == false {
		return nil, status.Errorf(codes.Internal, "更新用户文章数失败：%v", err)
	}
	return &pb.PostInfoDeleteResponse{
		Success: true,
	}, nil
}
