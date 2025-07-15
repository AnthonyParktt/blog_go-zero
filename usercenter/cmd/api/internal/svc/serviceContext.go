package svc

import (
	"go-zero_less/usercenter/cmd/api/internal/config"
	"go-zero_less/usercenter/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
