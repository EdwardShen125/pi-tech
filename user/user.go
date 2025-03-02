package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	protoc_gen_go "proto-gen-go"
	"user/internal/infrastructure/config"
	"user/internal/infrastructure/controller/rpc"
	"user/internal/svc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.MustSetUp()
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.UserRPCServer, func(grpcServer *grpc.Server) {
		protoc_gen_go.RegisterUserStdServiceServer(grpcServer, rpc.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.UserRPCServer.ListenOn)
	s.Start()
}
