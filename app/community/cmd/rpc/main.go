package main

import (
	"flag"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"
	"mall-go/app/community/cmd/pb"
	"mall-go/app/community/cmd/rpc/internal/config"
	"mall-go/app/community/cmd/rpc/internal/server"
	"mall-go/app/community/cmd/rpc/internal/svc"
	conf "mall-go/common/conf"
	_ "mall-go/pkg/di"
)

var configFile = flag.String("f", "etc/community.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)
	svc.Context = ctx
	// 创建server.CommunityServer结构体实例，主要是定义业务逻辑
	srv := server.NewCommunityServer(ctx)

	// 实际启动一个符合server.CommunityServer业务逻辑的grpc服务器
	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		// 函数参数srv签名本身是pb.CommunityServer接口，但是server.CommunityServer结构体实现了该接口，所以可以传入
		// Go编译器会自动匹配server.CommunityServer实现了pb.CommunityServer接口的方法
		pb.RegisterCommunityServer(g, srv)

	})

	defer s.Stop()

	s.Start()
}
