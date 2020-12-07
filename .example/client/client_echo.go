package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/katyusha-demos/protocol/pb"
	"github.com/gogf/katyusha/discovery"
	"github.com/gogf/katyusha/krpc"
	"golang.org/x/net/context"
	"time"
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

	echoClient := pb.NewEchoClient(conn)
	for i := 0; i < 500; i++ {
		res, err := echoClient.Say(context.Background(), &pb.SayReq{Content: "Hello"})
		if err != nil {
			g.Log().Error(err)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Second)
		g.Log().Print("Response:", res.Content)
	}
}
