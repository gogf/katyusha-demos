package service

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
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
	if err != nil {
		if gerror.Code(err) == int(user.ErrCode_UserNotFound) {
			g.Log().Debugf(`user "%s" not found, do auto signup`, passport)
			_, err = Client.User.SignUp(context.Background(), &user.SignUpReq{
				Passport: passport,
				Password: password,
				Nickname: passport,
			})
			if err != nil {
				g.Log().Fatal(err)
			}
			g.Log().Debugf(`user "%s" auto signup success, do auto login`, passport)
			signInRes, err = Client.User.SignIn(context.Background(), &user.SignInReq{
				Passport: passport,
				Password: password,
				Agent:    agent,
			})
			if err != nil {
				g.Log().Fatal(err)
			}
		} else {
			g.Log().Fatal(err)
		}
	}
	g.Log().Debugf(`user "%s" logins success, token: %s`, passport, signInRes.Token)
	for {
		time.Sleep(5 * time.Second)
		// 只要Session在使用则服务端会自动续期
		getSessionRes, err := Client.User.GetSession(context.Background(), &user.GetSessionReq{
			Token: signInRes.Token,
		})
		if err != nil {
			g.Log().Debugf(`auto keep login failed: %+v`, err)
		} else {
			g.Log().Debugf(
				`auto keep login, first login time: %s, agent: %s`,
				gtime.New(getSessionRes.LoginTime).String(),
				getSessionRes.LoginAgent,
			)
		}
	}
}
