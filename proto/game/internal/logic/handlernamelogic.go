package logic

import (
	"context"

	"game/internal/svc"
	"game/pb/protoc-gen-go"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlerNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandlerNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlerNameLogic {
	return &HandlerNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandlerNameLogic) HandlerName(in *protoc_gen_go.Request) (*protoc_gen_go.Response, error) {
	// todo: add your logic here and delete this line

	return &protoc_gen_go.Response{}, nil
}
