package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/katyusha-demos/app/service/demos-service/internal/dao"
	"github.com/gogf/katyusha-demos/app/service/demos-service/internal/model"
)

// 中间件管理服务
var User = new(serviceUser)

type serviceUser struct{}

// 用户注册
func (s *serviceUser) SignUp(ctx context.Context, r *model.ServiceUserSignUpReq) error {
	// 昵称为非必需参数，默认使用账号名称
	if r.Nickname == "" {
		r.Nickname = r.Passport
	}
	// 账号唯一性数据检查
	if !s.CheckPassport(ctx, r.Passport) {
		return gerror.New(fmt.Sprintf("账号 %s 已经存在", r.Passport))
	}
	// 昵称唯一性数据检查
	if !s.CheckNickName(ctx, r.Nickname) {
		return gerror.New(fmt.Sprintf("昵称 %s 已经存在", r.Nickname))
	}
	if _, err := dao.User.Ctx(ctx).Save(r); err != nil {
		return err
	}
	return nil
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *serviceUser) SignIn(ctx context.Context, req *model.ServiceUserSignInReq) (*model.User, error) {
	user, err := dao.User.Where(g.Slice{
		dao.User.Columns.Passport, req.Passport,
		dao.User.Columns.Password, req.Password,
	}).One()
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("账号或密码错误")
	}
	if err := Session.Set(ctx, &model.Session{
		User:       user,
		LoginTime:  gtime.Now(),
		LoginAgent: req.Agent,
	}); err != nil {
		return user, nil
	}
	Context.SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return user, nil
}

// 用户注销
func (s *serviceUser) SignOut(ctx context.Context) error {
	return Session.Remove(ctx)
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *serviceUser) CheckPassport(ctx context.Context, passport string) bool {
	if n, err := dao.User.Ctx(ctx).FindCount(dao.User.Columns.Passport, passport); err != nil {
		return false
	} else {
		return n > 0
	}
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *serviceUser) CheckNickName(ctx context.Context, nickname string) bool {
	if n, err := dao.User.Ctx(ctx).FindCount(dao.User.Columns.Nickname, nickname); err != nil {
		return false
	} else {
		return n > 0
	}
}

// 获得用户信息详情
func (s *serviceUser) GetUser(ctx context.Context, userId uint) (*model.User, error) {
	if user, err := dao.User.Ctx(ctx).FindOne(userId); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
