// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"mall-go/app/lottery/internal/config"
	"mall-go/app/lottery/internal/server"
	"mall-go/app/lottery/internal/svc"
)

import (
	_ "net/http/pprof"
)

// Injectors from wire.go:

// initApp init app application.
func initApp(c *config.Config) (*server.AppServer, error) {
	serviceContext := svc.NewServiceContext(c)
	lotteryServer := server.NewLotteryServer(serviceContext)
	httpServer := server.NewHttpServer(c, lotteryServer)
	rpcServer := server.NewGrpcServer(c, lotteryServer)
	appServer, err := server.NewApp(serviceContext, lotteryServer, httpServer, rpcServer)
	if err != nil {
		return nil, err
	}
	return appServer, nil
}
