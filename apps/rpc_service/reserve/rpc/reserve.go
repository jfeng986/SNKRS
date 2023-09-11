package main

import (
	"flag"
	"fmt"

	"snkrs/apps/rpc_service/reserve/rpc/internal/config"
	"snkrs/apps/rpc_service/reserve/rpc/internal/server"
	"snkrs/apps/rpc_service/reserve/rpc/internal/svc"
	"snkrs/apps/rpc_service/reserve/rpc/reserve"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/reserve.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		reserve.RegisterReserveServiceServer(grpcServer, server.NewReserveServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
