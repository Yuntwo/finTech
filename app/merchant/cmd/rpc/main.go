package main

import (
	"flag"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"
	"mall-go/app/merchant/cmd/pb"
	"mall-go/app/merchant/cmd/rpc/internal/config"
	"mall-go/app/merchant/cmd/rpc/internal/server"
	"mall-go/app/merchant/cmd/rpc/internal/svc"
	conf "mall-go/common/conf"
	_ "mall-go/pkg/di"
)

var configFile = flag.String("f", "etc/merchant.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)
	svc.Context = ctx

	srv := server.NewMerchantsServer(ctx)

	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		pb.RegisterMerchantsServer(g, srv)
	})

	defer s.Stop()

	s.Start()
}
