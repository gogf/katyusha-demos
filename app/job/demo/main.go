package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/katyusha-demos/app/job/demo/internal/service"
)

func main() {
	err := gcmd.BindHandleMap(map[string]func(){
		"echo": service.Echo.SayHiTimely,
		"user": func() {
			service.User.KeepLogin("john", "123")
		},
	})
	if err != nil {
		g.Log().Fatal(err)
	}
	if err = gcmd.AutoRun(); err != nil {
		g.Log().Fatal(err)
	}
}
