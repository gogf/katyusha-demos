package main

import (
	"github.com/gogf/katyusha-demos/app/service/echo/internal/api"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
	"github.com/gogf/katyusha/krpc"
)

func main() {
	c := krpc.Server.NewGrpcServerConfig()
	c.AppId = []string{}
	c.Options = append(
		c.Options,
	)
	s := krpc.Server.NewGrpcServer(c)
	echo.RegisterEchoServer(s.Server, api.Echo)
	s.Run()
}
