package main

import (
	"github.com/gogf/katyusha-demos/app/service/echo/internal/api"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
	"github.com/gogf/katyusha/krpc"
)

func main() {
	// 配置对象配置
	c := krpc.Server.NewGrpcServerConfig()
	// 指定服务名称
	c.AppId = echo.AppId

	// Server对象配置
	s := krpc.Server.NewGrpcServer(c)
	// 服务对象注册
	echo.RegisterEchoServer(s.Server, api.Echo)
	s.Run()
}
