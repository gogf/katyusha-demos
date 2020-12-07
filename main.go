package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/katyusha-demos/app/api"
	"github.com/gogf/katyusha-demos/protocol/pb"
	"github.com/gogf/katyusha/discovery"
	"github.com/gogf/katyusha/krpc"
)

func main() {
	genv.SetMap(g.MapStrStr{
		discovery.EnvKeyAppId:     `demo`,
		discovery.EnvKeyMetaData:  `{"weight":100}`,
		discovery.EnvKeyEndpoints: "127.0.0.1:2379",
	})
	s := krpc.NewGrpcServer()
	pb.RegisterEchoServer(s.Server, api.Echo)
	pb.RegisterUserServer(s.Server, api.User)
	s.Run()
}
