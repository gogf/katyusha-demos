package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
	"github.com/gogf/katyusha-demos/protobuf/demos"
	"github.com/gogf/katyusha/krpc"
	"golang.org/x/net/context"
	"time"
)

var Echo = echoApi{}

type echoApi struct{}

func (a *echoApi) Say(r *ghttp.Request) {
	sayRes, err := service.Echo.Say(r.Context(), &echo.SayReq{
		Content: content,
	})
	if err != nil {
		return "", err
	}
	return sayRes.Content, nil
}
