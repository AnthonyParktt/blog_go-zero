package login

import (
	"context"
	"go-zero_less/usercenter/cmd/api/internal/svc"
	"go-zero_less/usercenter/cmd/api/internal/types"
	"go-zero_less/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
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
	// user, err := l.svcCtx.UserModel.FindOneByUname(l.ctx, req.Username)
	// if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
	// 	return nil, err
	// }
	// if user == nil {
	// 	return nil, errors.New("user not found")
	// }
	// // 密码校验
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	// if err != nil {
	// 	return nil, errors.New("用户名密码错误")
	// }
	// token, err := utils.GenerateToken(uint(user.Id), "123456")
	// if err != nil {
	// 	return nil, errors.New("token生成错误")
	// }
	loginReq := &usercenter.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}
	loginResp, err := l.svcCtx.UserRpcClient.Login(l.ctx, loginReq)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Id:    loginResp.Id,
		Name:  loginResp.Username,
		Token: loginResp.Token,
	}, nil
}
