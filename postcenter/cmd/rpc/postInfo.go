package main

import (
	"flag"
	"fmt"

	"go-zero_less/postcenter/cmd/rpc/internal/config"
	commentsinfoServer "go-zero_less/postcenter/cmd/rpc/internal/server/commentsinfo"
	postinfoServer "go-zero_less/postcenter/cmd/rpc/internal/server/postinfo"
	"go-zero_less/postcenter/cmd/rpc/internal/svc"
	"go-zero_less/postcenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/postInfo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterPostInfoServer(grpcServer, postinfoServer.NewPostInfoServer(ctx))
		pb.RegisterCommentsInfoServer(grpcServer, commentsinfoServer.NewCommentsInfoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
