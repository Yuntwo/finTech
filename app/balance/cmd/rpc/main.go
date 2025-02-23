package main

import (
	"flag"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"
	"mall-go/app/balance/cmd/pb"
	"mall-go/app/balance/cmd/rpc/internal/config"
	"mall-go/app/balance/cmd/rpc/internal/server"
	"mall-go/app/balance/cmd/rpc/internal/svc"
	conf "mall-go/common/conf"
	_ "mall-go/pkg/di"
)

var configFile = flag.String("f", "etc/balance.yaml", "the config file")

func main() {
	flag.Parse()

	// 1.加载配置，初始化 ServiceContext
	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)
	svc.Context = ctx
	srv := server.NewBalanceServer(ctx)
	// 2.注册 gRPC 服务和接口
	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		pb.RegisterBalanceServer(g, srv)
	})
	// 3.启动 gRPC 服务监听，处理请求
	defer s.Stop()

	s.Start()
}
