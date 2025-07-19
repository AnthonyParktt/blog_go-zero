package svc

import (
	"go-zero_less/usercenter/cmd/api/internal/config"
	"go-zero_less/usercenter/cmd/rpc/usercenter"
	"go-zero_less/usercenter/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UsersModel
	UserRpcClient usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource)),
		UserRpcClient: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpc)),
	}
}
