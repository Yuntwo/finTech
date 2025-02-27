package server

import (
	lottery "mall-go/app/marketing/api/lottery/v1"
	"net/http"

	"context"

	"mall-go/app/marketing/internal/svc"

	"github.com/google/wire"
	"github.com/mix-plus/go-mixplus/mrpc"
)

// ProviderSet is server providers.
// ProviderSet主要是方便组合多个提供者(即函数)放在一起，也可以不这样包装直接一个一个传
// 还可以使用wire.Struct将某个结构体中的字段标注为依赖项
// 注意区别这三个server，前两者是具体通信协议的server实例，最后一个是业务逻辑的而非实际启动接口的server实例
var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer, NewLotteryServer)

type AppServer struct {
	SvcCtx     *svc.ServiceContext
	HttpServer *http.Server
	GrpcServer *mrpc.RpcServer

	HelloService *LotteryServer
}

func NewApp(svcCtx *svc.ServiceContext, helloService *LotteryServer, hs *http.Server, gs *mrpc.RpcServer) (*AppServer, error) {
	return &AppServer{
		SvcCtx:       svcCtx,
		HelloService: helloService,
		HttpServer:   hs,
		GrpcServer:   gs,
	}, nil
}

func (a *AppServer) Run() {

	go func() {
		err := a.HttpServer.ListenAndServe()
		if err != nil {
			return
		}
	}()

	a.GrpcServer.Start()

	defer a.GrpcServer.Stop()
}

type LotteryServer struct {
	// 实际上继承了Lottery
	lottery.UnimplementedLotteryServer

	svcCtx *svc.ServiceContext
}

func NewLotteryServer(ctx *svc.ServiceContext) *LotteryServer {
	return &LotteryServer{
		svcCtx: ctx,
	}
}

func (l *LotteryServer) FindLottery(ctx context.Context, req *lottery.FindLotteryReq) (*lottery.LotteryResp, error) {
	return &lottery.LotteryResp{
		Id:          req.Id,
		Name:        "x",
		Description: "x",
	}, nil
}

func (l *LotteryServer) CreateLottery(ctx context.Context, req *lottery.CreateLotteryReq) (*lottery.LotteryResp, error) {
	return &lottery.LotteryResp{
		Id:          1,
		Name:        req.Name,
		Description: req.Description,
	}, nil
}
