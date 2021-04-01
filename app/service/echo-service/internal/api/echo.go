package api

import (
	"context"
	"fmt"
	"github.com/gogf/katyusha-demos/app/service/demos-service/protobuf/demos"
)

// 回显服务
var Echo = &echoApi{}

type echoApi struct{}

func (a *echoApi) Say(ctx context.Context, r *demos.SayReq) (*demos.SayRes, error) {
	text := fmt.Sprintf(`> %s`, r.Content)
	return &demos.SayRes{Content: text}, nil
}
