package logic

import (
	"context"
	"errors"

	"go-zero_less/usercenter/cmd/rpc/internal/svc"
	"go-zero_less/usercenter/cmd/rpc/pb"
	"go-zero_less/usercenter/model"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersLogic) GetUsers(in *pb.UserInfosReq) (*pb.UserInfosResp, error) {
	users, err := l.svcCtx.UserModel.FindUsers(l.ctx, model.UserQueryParams{
		Username: in.Username,
		Email:    in.Email,
		Status:   in.Status,
	})
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, status.Errorf(codes.Internal, "用户查询失败", err)
	}
	//将users转换为pb.UserInfosResp
	var resp []*pb.UserInfo
	for _, user := range users {
		resp = append(resp, &pb.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Status:   user.Status,
		})
	}

	return &pb.UserInfosResp{
		UserInfos: resp,
	}, nil
}
