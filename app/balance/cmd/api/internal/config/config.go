package config

import (
	"github.com/mix-plus/go-mixplus/mrpc"
	"mall-go/common/conf"
)

type Config struct {
	conf.ApiConf `mapstructure:",squash"`
	conf.JwtAuth
	BalanceRpcConf mrpc.RpcClientConf
}
