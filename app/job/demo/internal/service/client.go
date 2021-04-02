package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
	"github.com/gogf/katyusha-demos/app/service/user/protobuf/user"
	"github.com/gogf/katyusha/krpc"
)

// GRPC客户端对象统一管理
var Client = clientService{}

type clientService struct {
	Echo echo.EchoClient
	User user.UserClient
}

// 模块初始化时创建GRPC客户端对象。
// 注意：GRPC客户端对象只有在真正通信时才会创建连接。
func init() {
	echoConn, err := krpc.Client.NewGrpcClientConn(echo.AppId)
	if err != nil {
		g.Log().Fatal(err)
	}
	userConn, err := krpc.Client.NewGrpcClientConn(user.AppId)
	if err != nil {
		g.Log().Fatal(err)
	}
	Client = clientService{
		Echo: echo.NewEchoClient(echoConn),
		User: user.NewUserClient(userConn),
	}
}
