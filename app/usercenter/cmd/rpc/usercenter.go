package main

import (
	"flag"
	"fmt"

	"trytry/app/usercenter/cmd/rpc/internal/config"
	"trytry/app/usercenter/cmd/rpc/internal/server"
	"trytry/app/usercenter/cmd/rpc/internal/svc"
	"trytry/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUsercenterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//拦截器 在服务端不常用
	//s.AddUnaryInterceptors()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
