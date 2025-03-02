package rpc

import (
	"context"
	protoc_gen_go "proto-gen-go/user"

	"user/internal/svc"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	protoc_gen_go.UnimplementedUserStdServiceServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Ping(ctx context.Context, in *protoc_gen_go.Request) (*protoc_gen_go.Response, error) {
	user, err := s.svcCtx.UserAppSvc.RegisterUser(s.ctx, "xxx", "xxx", "xxx", "xxx")
	return nil, err
}
