package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
	"time"
)

// 回显业务逻辑封装
var Echo = echoService{}

type echoService struct{}

// 打个招呼
func (s *echoService) Say(ctx context.Context, content string) (reply string, err error) {
	sayRes, err := Client.Echo.Say(ctx, &echo.SayReq{
		Content: content,
	})
	if err != nil {
		return "", err
	}
	return sayRes.Content, nil
}

// 每隔1秒打一次招呼
func (s *echoService) SayHiTimely() {
	for {
		content := "hi"
		g.Log().Print("  say:", content)
		reply, err := s.Say(context.Background(), content)
		if err != nil {
			g.Log().Print(err)
		} else {
			g.Log().Print("reply:", reply)
		}
		time.Sleep(time.Second)
	}
}
