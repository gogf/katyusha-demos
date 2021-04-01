package main

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/katyusha-demos/app/job/demo/internal/service"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
)

func main() {
	r, err := service.Client.Echo.Say(context.Background(), &echo.SayReq{
		Content: "hello",
	})
	if err != nil {
		g.Log().Fatal(err)
	}
	g.Log().Print(r)
}
