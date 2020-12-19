package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/katyusha-demos/protobuf/demos"
	"github.com/gogf/katyusha/krpc"
	"golang.org/x/net/context"
	"time"
)

func main() {
	conn, err := krpc.Client.NewGrpcClientConn("demos")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	echoClient := demos.NewEchoClient(conn)
	for i := 0; i < 500; i++ {
		res, err := echoClient.Say(context.Background(), &demos.SayReq{Content: "Hello"})
		if err != nil {
			g.Log().Error(err)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Second)
		g.Log().Print("Response:", res.Content)
	}
}
