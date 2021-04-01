package main

import (
	"github.com/gogf/katyusha-demos/app/service/user/internal/api"
	"github.com/gogf/katyusha-demos/app/service/user/internal/service"
	"github.com/gogf/katyusha-demos/app/service/user/protobuf/user"
	"github.com/gogf/katyusha/krpc"
)

func main() {
	c := krpc.Server.NewGrpcServerConfig()
	c.Options = append(
		c.Options,
		krpc.Server.ChainUnary(
			service.Interceptor.UnaryCtx,
			service.Interceptor.UnaryValidate,
		),
	)
	s := krpc.Server.NewGrpcServer(c)
	user.RegisterUserServer(s.Server, api.User)
	s.Run()
}
