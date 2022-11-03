package config

import "github.com/mix-plus/core/conf"

type Config struct {
	conf.ApiConf `mapstructure:",squash"`
}
