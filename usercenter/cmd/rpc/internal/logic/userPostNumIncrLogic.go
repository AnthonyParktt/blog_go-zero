package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-zero_less/usercenter/cmd/rpc/internal/svc"
	"go-zero_less/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPostNumIncrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPostNumIncrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPostNumIncrLogic {
	return &UserPostNumIncrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPostNumIncrLogic) UserPostNumIncr(in *pb.UserPostNumReq) (*pb.ResultBool, error) {
	err := l.svcCtx.UserModel.AddUserPostNum(l.ctx, in.Id, in.PostNum)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "用户文章数量更新失败：%v", err)
	}
	return &pb.ResultBool{
		Result: true,
	}, nil
}
