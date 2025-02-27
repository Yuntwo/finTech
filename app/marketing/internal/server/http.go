package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	lottery "mall-go/app/marketing/api/lottery/v1"
	"mall-go/app/marketing/internal/config"
)

func NewHttpServer(c *config.Config, srv *LotteryServer) *http.Server {
	httpServer := &http.Server{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	err := lottery.RegisterLotteryHandlerServer(ctx, mux, srv)
	if err != nil {
		panic(err)
	}
	httpServer.Addr = c.ApiConf.Addr
	httpServer.Handler = mux

	return httpServer
}
