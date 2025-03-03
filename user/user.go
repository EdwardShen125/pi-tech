package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net/http"
	protoc_gen_go "proto-gen-go"
	"user/internal/infrastructure/config"
	"user/internal/infrastructure/controller/rpc"
	"user/internal/infrastructure/mqs"
	"user/internal/svc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	// conf
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.MustSetUp()
	ctx := svc.NewServiceContext(c)

	// service group
	sg := service.NewServiceGroup()
	defer sg.Stop()

	// consumers
	for _, mq := range mqs.Consumers(c, context.Background(), ctx) {
		sg.Add(mq)
	}

	// rpc server
	s := zrpc.MustNewServer(c.UserRPCServer, func(grpcServer *grpc.Server) {
		protoc_gen_go.RegisterUserStdServiceServer(grpcServer, rpc.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	fmt.Printf("Starting rpc server at %s...\n", c.UserRPCServer.ListenOn)
	sg.Add(s)

	rest.MustNewServer(rest.RestConf{})

	// start all service
	sg.Start()
}
