package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/katyusha-demos/app/api"
	"github.com/gogf/katyusha-demos/app/service"
	"github.com/gogf/katyusha-demos/protobuf/demos"
	"github.com/gogf/katyusha/discovery"
	"github.com/gogf/katyusha/krpc"
)

func main() {
	genv.SetMap(g.MapStrStr{
		discovery.EnvKeyAppId:     `demos`,
		discovery.EnvKeyEndpoints: "127.0.0.1:2379",
	})
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
