package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/katyusha-demos/app/service/user/internal/dao"
	"github.com/gogf/katyusha-demos/app/service/user/internal/model"
	"github.com/gogf/katyusha-demos/app/service/user/protobuf/user"
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
	if result, err := s.CheckPassport(ctx, r.Passport); err != nil {
		return err
	} else if !result {
		return gerror.New(fmt.Sprintf("账号 %s 已经存在", r.Passport))
	}
	// 昵称唯一性数据检查
	if result, err := s.CheckNickName(ctx, r.Nickname); err != nil {
		return err
	} else if !result {
		return gerror.New(fmt.Sprintf("昵称 %s 已经存在", r.Nickname))
	}
	if _, err := dao.User.Ctx(ctx).Save(r); err != nil {
		return err
	}
	return nil
}

// 用户登录，成功返回用户信息，否则返回nil。
// 为了使用到错误码演示，这里会判断并返回登录错误原因。
func (s *serviceUser) SignIn(ctx context.Context, req *model.ServiceUserSignInReq) (*model.User, error) {
	entity, err := dao.User.Where(dao.User.Columns.Passport, req.Passport).One()
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, gerror.NewCodef(int(user.ErrCode_UserNotFound), `账号"%s"不存在`, req.Passport)
	}
	if entity.Password != req.Password {
		return nil, gerror.NewCode(int(user.ErrCode_PasswordIncorrect), `密码不正确`)
	}
	if err := Session.Set(ctx, &model.Session{
		User:       entity,
		LoginTime:  gtime.Now(),
		LoginAgent: req.Agent,
	}); err != nil {
		return entity, nil
	}
	Context.SetUser(ctx, &model.ContextUser{
		Id:       entity.Id,
		Passport: entity.Passport,
		Nickname: entity.Nickname,
	})
	return entity, nil
}

// 用户注销
func (s *serviceUser) SignOut(ctx context.Context) error {
	return Session.Remove(ctx)
}

// 检查账号是否可用(目前仅检查唯一性),存在返回false,否则true
func (s *serviceUser) CheckPassport(ctx context.Context, passport string) (bool, error) {
	if n, err := dao.User.Ctx(ctx).FindCount(dao.User.Columns.Passport, passport); err != nil {
		return false, err
	} else {
		return n == 0, nil
	}
}

// 检查昵称是否可用(目前仅检查唯一性),存在返回false,否则true
func (s *serviceUser) CheckNickName(ctx context.Context, nickname string) (bool, error) {
	if n, err := dao.User.Ctx(ctx).FindCount(dao.User.Columns.Nickname, nickname); err != nil {
		return false, err
	} else {
		return n == 0, nil
	}
}

// 获得用户信息详情
func (s *serviceUser) GetUser(ctx context.Context, userId uint) (*model.User, error) {
	if entity, err := dao.User.Ctx(ctx).FindOne(userId); err != nil {
		return nil, err
	} else {
		return entity, nil
	}
}
