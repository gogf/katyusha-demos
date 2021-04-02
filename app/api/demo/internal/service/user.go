package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/katyusha-demos/app/api/demo/internal/model"
	"github.com/gogf/katyusha-demos/app/service/user/protobuf/user"
)

// 中间件管理服务
var User = userService{}

type userService struct{}

// 用户注册
func (s *userService) SignUp(ctx context.Context, r *model.UserServiceSignUpReq) error {
	// 昵称为非必需参数，默认使用账号名称
	if r.Nickname == "" {
		r.Nickname = r.Passport
	}
	var signUpReq *user.SignUpReq
	if err := gconv.Struct(r, &signUpReq); err != nil {
		return err
	}
	if _, err := Client.User.SignUp(ctx, signUpReq); err != nil {
		return err
	}
	return nil
}

// 判断用户是否已经登录
func (s *userService) IsSignedIn(ctx context.Context, token string) (ok bool, err error) {
	_, err = Client.User.GetSession(ctx, &user.GetSessionReq{
		Token: token,
	})
	if err != nil {
		if gerror.Code(err) == int(user.ErrCode_UserNotLogin) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *userService) SignIn(ctx context.Context, r *model.UserServiceSignInReq) (*model.UserServiceSignInRes, error) {
	var signInReq *user.SignInReq
	if err := gconv.Struct(r, &signInReq); err != nil {
		return nil, err
	}
	signInRes, err := Client.User.SignIn(ctx, signInReq)
	if err != nil {
		return nil, err
	}
	return &model.UserServiceSignInRes{Token: signInRes.Token}, nil
}

// 用户注销
func (s *userService) SignOut(ctx context.Context) error {
	return Session.RemoveUser(ctx)
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckPassport(ctx context.Context, passport string) (ok bool, err error) {
	checkPassportRes, err := Client.User.CheckPassport(ctx, &user.CheckPassportReq{
		Passport: passport,
	})
	if err != nil {
		return false, err
	}
	return checkPassportRes.Ok, nil
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckNickName(ctx context.Context, nickname string) (ok bool, err error) {
	checkNickNameRes, err := Client.User.CheckNickName(ctx, &user.CheckNickNameReq{
		Nickname: nickname,
	})
	if err != nil {
		return false, err
	}
	return checkNickNameRes.Ok, nil
}

// 获得用户信息详情
func (s *userService) GetProfile(ctx context.Context) *model.User {
	return Session.GetUser(ctx)
}
