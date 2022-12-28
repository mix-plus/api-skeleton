package svc

import (
	hello "github.com/mix-plus/api-skeleton/api"
	"github.com/mix-plus/api-skeleton/internal/config"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config

	// Register RPC
	HelloRpc hello.HelloClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//HelloRpc: hello.NewHelloClient(mrpc.MustNewClient(c.HelloRpcConf).Conn()),
	}
}
