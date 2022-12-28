package config

import (
	"github.com/go-ll/mrpc"
	"github.com/mix-plus/core/conf"
)

type Config struct {
	conf.ApiConf `mapstructure:",squash"`

	HelloRpcConf mrpc.RpcClientConf
}
