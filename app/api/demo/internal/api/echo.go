package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/katyusha-demos/app/api/demo/internal/model"
	"github.com/gogf/katyusha-demos/app/api/demo/internal/service"
	"github.com/gogf/katyusha-demos/library/response"
)

// 回显服务
var Echo = echoApi{}

type echoApi struct{}

func (a *echoApi) Say(r *ghttp.Request) {
	var (
		apiReq *model.EchoApiSayReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	reply, err := service.Echo.Say(r.Context(), apiReq.Content)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, reply)
}
