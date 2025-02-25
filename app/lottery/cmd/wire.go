// go:build wireinject和+build wireinject是构建标记，用于确保这个文件只有在构建时[显式]使用wireinject构建标记时才被包含在构建过程中
// 就是go build　--tags(选项) wireinject的时候才会编译这个文件，在goland中也会提示"'wire.go' is ignored by the build tool because of the customflag"，这个是正常的
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
// 在Makefile中没看到生成对应的文件，一般来说还是手动生成的，否则编译也要报错
// 按理说在wire.go文件中也添加//go:generate go run github.com/google/wire/cmd/wire这个注释方便点，但是生成文件中有这个注释(其实就多余了)算双重保险吧，不太理解
package main

import (
	"mall-go/app/lottery/internal/config"
	"mall-go/app/lottery/internal/server"
	"mall-go/app/lottery/internal/svc"

	"github.com/google/wire"
)

// initApp init app application.
func initApp(c *config.Config) (*server.AppServer, error) {
	// panic是内建函数，当函数返回错误时，程序会中断执行，打印错误信息
	// wire.Build用来指定Wire如何构建依赖关系图。通过将多个ProviderSet和server.NewApp传递给wire.Build，Wire会生成一个可以自动注入所有依赖的代码
	// 这个函数会根据函数前面自动解析依赖关系，比如出现循环依赖的话会检查出报错
	// 传递的都是函数签名，底层处理的是各个函数的参数和返回值依赖关系
	// 生成的InitApp是核心函数，主要是根据wire.go中的定义手动实例化了每个服务并将它们注入到最终的appServer中
	panic(wire.Build(svc.ProviderSet, server.ProviderSet, server.NewApp))
}
