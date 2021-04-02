package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/katyusha-demos/app/api/demo/internal/api"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALLMap(g.Map{
			"/echo": api.Echo,
		})
	})
	s.Run()
}
