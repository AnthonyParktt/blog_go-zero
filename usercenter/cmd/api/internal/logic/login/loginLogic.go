package login

import (
	"context"
	"errors"
	"go-zero_less/pkg/utils"
	"go-zero_less/usercenter/cmd/api/internal/svc"
	"go-zero_less/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUname(l.ctx, req.Username)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	// 密码校验
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("用户名密码错误")
	}
	token, err := utils.GenerateToken(uint(user.Id), "123456")
	if err != nil {
		return nil, errors.New("token生成错误")
	}
	return &types.LoginResp{
		Id:    user.Id,
		Name:  user.Username,
		Token: token,
	}, nil
}
