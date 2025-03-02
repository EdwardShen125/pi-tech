// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: game_service.proto

package server

import (
	"context"

	"game/internal/logic"
	"game/internal/svc"
	"game/pb/protoc-gen-go"
)

type GameServiceServer struct {
	svcCtx *svc.ServiceContext
	protoc_gen_go.UnimplementedGameServiceServer
}

func NewGameServiceServer(svcCtx *svc.ServiceContext) *GameServiceServer {
	return &GameServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *GameServiceServer) HandlerName(ctx context.Context, in *protoc_gen_go.Request) (*protoc_gen_go.Response, error) {
	l := logic.NewHandlerNameLogic(ctx, s.svcCtx)
	return l.HandlerName(in)
}
