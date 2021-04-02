package model

type EchoApiSayReq struct {
	Content string `v:"required#内容不能为空"`
}
