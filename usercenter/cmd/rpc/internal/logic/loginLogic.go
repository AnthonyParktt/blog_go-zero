package logic

import (
	"context"
	"errors"

	"go-zero_less/pkg/utils"
	"go-zero_less/usercenter/cmd/rpc/internal/svc"
	"go-zero_less/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	user, err := l.svcCtx.UserModel.FindUserByName(l.ctx, in.Username)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, err
	}
	if user == nil {
		return nil, status.Error(codes.Unauthenticated, "用户名密码错误")
	}
	// 密码校验
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "用户名密码错误")
	}
	token, err := utils.GenerateToken(uint(user.Id), "123456")
	if err != nil {
		return nil, errors.New("token生成错误")
	}

	return &pb.LoginResp{
		Id:       user.Id,
		Username: in.Username,
		Token:    token,
		ExpireAt: "",
	}, nil
}
