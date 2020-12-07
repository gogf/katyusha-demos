package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/katyusha-demos/protocol/pb"
	"github.com/gogf/katyusha/discovery"
	"github.com/gogf/katyusha/krpc"
	"golang.org/x/net/context"
)

func main() {
	genv.SetMap(g.MapStrStr{
		discovery.EnvKeyEndpoints: "127.0.0.1:2379",
	})

	conn, err := krpc.NewGrpcClientConn("demo")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	res, err := client.SignUp(context.Background(), &pb.SignUpReq{
		Passport:  "john2",
		Password:  "123",
		Password2: "123",
		Nickname:  "JohnGuo2",
	})
	g.Log().Print(res, err)
}
