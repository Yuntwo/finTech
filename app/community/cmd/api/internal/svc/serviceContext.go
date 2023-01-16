package svc

import (
	"mall-go/app/community/cmd/api/internal/config"
	"mall-go/app/community/cmd/pb"
	"mall-go/pkg/jwtx"

	"github.com/mix-plus/go-mixplus/mrpc"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config
	Jwt    *jwtx.Jwt

	CommunityRpc pb.CommunityClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Jwt:          jwtx.NewJwt(c.JwtAuth),
		CommunityRpc: pb.NewCommunityClient(mrpc.MustNewClient(c.CommunityRpcConf).Conn()),
	}
}
