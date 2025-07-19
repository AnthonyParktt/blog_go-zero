package logic

import (
	"context"
	"errors"

	"go-zero_less/usercenter/cmd/rpc/internal/svc"
	"go-zero_less/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.UserInfoOneReq) (*pb.UserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOneUser(l.ctx, in.Id, in.Username)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, status.Errorf(codes.Internal, "查询用户失败：%v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return &pb.UserInfoResp{
		UserInfo: &pb.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}
