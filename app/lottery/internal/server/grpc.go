package server

import (
	lottery "mall-go/app/lottery/api/lottery/v1"
	"mall-go/app/lottery/internal/config"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"
)

func NewGrpcServer(c *config.Config, srv *LotteryServer) *mrpc.RpcServer {

	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		// grpc register
		lottery.RegisterLotteryServer(g, srv)
	})

	return s
}
