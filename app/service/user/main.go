package main

import (
	"github.com/gogf/katyusha-demos/app/service/user/internal/api"
	"github.com/gogf/katyusha-demos/app/service/user/internal/service"
	"github.com/gogf/katyusha-demos/app/service/user/protobuf/user"
	"github.com/gogf/katyusha/krpc"
)

func main() {
	// 配置对象配置
	c := krpc.Server.NewGrpcServerConfig()
	// 指定服务名称
	c.AppId = user.AppId
	// 拦截器注册
	c.Options = append(
		c.Options,
		krpc.Server.ChainUnary(
			service.Interceptor.UnaryCtx,
			service.Interceptor.UnaryValidate,
		),
	)

	// Server对象配置
	s := krpc.Server.NewGrpcServer(c)
	// 服务对象注册
	user.RegisterUserServer(s.Server, api.User)
	s.Run()
}
