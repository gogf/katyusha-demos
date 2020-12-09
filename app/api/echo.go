package api

import (
	"context"
	"fmt"
	"github.com/gogf/katyusha-demos/protobuf/demos"
)

// 回显服务
var Echo = &apiEcho{}

type apiEcho struct{}

func (s *apiEcho) Say(ctx context.Context, r *demos.SayReq) (*demos.SayRes, error) {
	text := fmt.Sprintf(`> %s`, r.Content)
	return &demos.SayRes{Content: text}, nil
}
