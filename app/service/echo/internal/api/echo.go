package api

import (
	"context"
	"fmt"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
)

// 回显服务
var Echo = &echoApi{}

type echoApi struct{}

func (a *echoApi) Say(ctx context.Context, r *echo.SayReq) (*echo.SayRes, error) {
	text := fmt.Sprintf(`> %s`, r.Content)
	return &echo.SayRes{Content: text}, nil
}
