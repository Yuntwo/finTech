package viper

import (
	"fmt"

	config "github.com/cexll/mall-go/config"
	"github.com/mix-go/xcli/argv"
	"github.com/spf13/viper"
)

func init() {
	// Conf support JSON, TOML, YAML, HCL, INI, envfile
	viper.SetConfigFile(fmt.Sprintf("%s/../conf/config.yml", argv.Program().Dir))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config.Config); err != nil {
		panic(err)
	}
}
