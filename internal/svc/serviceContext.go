package svc

import "github.com/mix-plus/api-skeleton/internal/config"

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
