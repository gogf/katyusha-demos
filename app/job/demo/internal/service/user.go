package service

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/katyusha-demos/app/service/user/protobuf/user"
	"time"
)

// 用户业务逻辑封装
var User = userService{}

type userService struct{}

// 保持用户在线状态（session不退出）
func (s *userService) KeepLogin(passport, password string) {
	agent := "demo-job"
	signInRes, err := Client.User.SignIn(context.Background(), &user.SignInReq{
		Passport: passport,
		Password: password,
		Agent:    agent,
	})
	if err != nil && gerror.Code(err) == int(user.ErrCode_UserNotFound) {
		g.Log().Printf(`user "%s" not found, do auto signup`, passport)
		_, err = Client.User.SignUp(context.Background(), &user.SignUpReq{
			Passport: passport,
			Password: password,
			Nickname: passport,
		})
		if err != nil {
			g.Log().Fatal(err)
		}
		g.Log().Printf(`user "%s" do auto signup success, do auto login`, passport)
		signInRes, err = Client.User.SignIn(context.Background(), &user.SignInReq{
			Passport: passport,
			Password: password,
			Agent:    agent,
		})
		if err != nil {
			g.Log().Fatal(err)
		}
	}
	g.Log().Printf(`user "%s" logins success, token: %s`, signInRes.Token)
	for {
		time.Sleep(time.Second)
		getSessionRes, err := Client.User.GetSession(context.Background(), &user.GetSessionReq{
			Token: signInRes.Token,
		})
		if err != nil {
			g.Log().Printf(`auto keep login failed: %+v`, err)
		} else {
			g.Log().Printf(
				`auto keep login success, login time: %s, login agent: %s`,
				getSessionRes.LoginTime,
				getSessionRes.LoginAgent,
			)
		}
	}
}
