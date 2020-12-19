package main

import (
	"github.com/gogf/katyusha-demos/app/api"
	"github.com/gogf/katyusha-demos/app/service"
	"github.com/gogf/katyusha-demos/protobuf/demos"
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
	demos.RegisterEchoServer(s.Server, api.Echo)
	demos.RegisterUserServer(s.Server, api.User)
	s.Run()
}
