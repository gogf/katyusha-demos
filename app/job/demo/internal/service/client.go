package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
	"github.com/gogf/katyusha-demos/app/service/user/protobuf/user"
	"github.com/gogf/katyusha/krpc"
)

var Client = clientService{}

type clientService struct {
	Echo echo.EchoClient
	User user.UserClient
}

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
