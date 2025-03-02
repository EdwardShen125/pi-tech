package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/app"
	"user/internal/domain/service"
	"user/internal/infrastructure/config"
	"user/internal/infrastructure/repository/event"
	"user/internal/infrastructure/repository/user"
)

type ServiceContext struct {
	Config     config.Config
	UserAppSvc *app.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化sql connection
	sqlConn := sqlx.NewMysql(c.MysqlDSN)
	// 初始化仓储
	userRepo := user.NewUserModel(sqlConn)
	eventRepo := event.NewEventModel(sqlConn)

	// 初始化领域服务
	userDomainSvc := service.NewUserDomainService()

	// 初始化应用层服务
	userAppSvc := app.NewUserService(
		sqlConn,
		userRepo,
		userDomainSvc,
		eventRepo,
	)

	return &ServiceContext{
		Config:     c,
		UserAppSvc: userAppSvc, // 显示传递应用层服务
	}
}
