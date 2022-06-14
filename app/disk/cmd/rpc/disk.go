package main

import (
	"flag"
	"fmt"

	"cloud-disk/app/disk/cmd/rpc/internal/config"
	"cloud-disk/app/disk/cmd/rpc/internal/server"
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/disk.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewStoreServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterStoreServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
