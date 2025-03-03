package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"user/internal/domain/entity"
	"user/internal/svc"
)

type UserRegistered struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegistered(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegistered {
	return &UserRegistered{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegistered) Consume(ctx context.Context, key, val string) error {
	logx.Infof("UserRegistered key :%s , val :%s", key, val)
	user := &entity.UserRegisteredEvent{}
	err := json.Unmarshal([]byte(val), user)
	if err != nil {
		return err
	}

	return l.svcCtx.UserAppSvc.HandleRegisterUser(ctx, user)
}
