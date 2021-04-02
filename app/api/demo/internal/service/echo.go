package service

import (
	"context"
	"github.com/gogf/katyusha-demos/app/service/echo/protobuf/echo"
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
