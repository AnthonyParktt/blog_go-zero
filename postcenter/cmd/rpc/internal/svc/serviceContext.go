package svc

import (
	"go-zero_less/postcenter/cmd/rpc/internal/config"
	"go-zero_less/postcenter/model"
	"go-zero_less/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	PostModel        model.PostsModel
	CommentsModel    model.CommentsModel
	UserCenterClient usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		PostModel:        model.NewPostsModel(sqlx.NewMysql(c.DB.DataSource)),
		CommentsModel:    model.NewCommentsModel(sqlx.NewMysql(c.DB.DataSource)),
		UserCenterClient: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcClient)),
	}
}
