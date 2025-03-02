package config

import (
	"errors"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	Etcd          discov.EtcdConf
	UserRPCServer zrpc.RpcServerConf
	GameRPCClient zrpc.RpcClientConf
	MysqlDSN      string
}

// Validate 校验配置
func (c Config) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("name 不能为空")
	}

	// ...
	return nil
}
