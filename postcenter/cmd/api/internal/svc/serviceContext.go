package svc

import (
	"go-zero_less/postcenter/cmd/api/internal/config"
	"go-zero_less/postcenter/cmd/rpc/client/postinfo"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	PostRpcClient postinfo.PostInfo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		PostRpcClient: postinfo.NewPostInfo(zrpc.MustNewClient(c.PostRpcClient)),
	}
}
