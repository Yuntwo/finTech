package svc

import (
	"mall-go/app/marketing/internal/config"

	"github.com/google/wire"
)

// ProviderSet is server providers.
// 这样直接声明都是包级变量
var ProviderSet = wire.NewSet(NewServiceContext)

type ServiceContext struct {
	Config *config.Config
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
