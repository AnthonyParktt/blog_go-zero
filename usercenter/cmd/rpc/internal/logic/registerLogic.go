package logic

import (
	"context"

	"go-zero_less/usercenter/cmd/rpc/internal/svc"
	"go-zero_less/usercenter/cmd/rpc/pb"
	"go-zero_less/usercenter/model"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	userModel := model.Users{
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
	}
	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userModel.Password = string(hash)
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &userModel)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Internal, "注册失败")
	}
	return &pb.RegisterResp{
		Id: uint64(id),
	}, nil
}
