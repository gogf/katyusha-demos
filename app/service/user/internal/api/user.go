package api

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/katyusha-demos/app/service/user/internal/model"
	"github.com/gogf/katyusha-demos/app/service/user/internal/service"
	"github.com/gogf/katyusha-demos/app/service/user/protobuf/user"
)

// 用户服务
var User = &userApi{}

type userApi struct{}

// 注册
func (s *userApi) SignUp(ctx context.Context, req *user.SignUpReq) (*user.NullRes, error) {
	var (
		res        user.NullRes
		serviceReq *model.ServiceUserSignUpReq
	)
	if err := gconv.Struct(req, &serviceReq); err != nil {
		return nil, gerror.Current(err)
	}
	err := service.User.SignUp(ctx, serviceReq)
	return &res, err
}

// 登录
func (s *userApi) SignIn(ctx context.Context, req *user.SignInReq) (*user.SignInRes, error) {
	var (
		res        = user.SignInRes{}
		serviceReq *model.ServiceUserSignInReq
	)
	if err := gconv.Struct(req, &serviceReq); err != nil {
		return nil, err
	}
	user, err := service.User.SignIn(ctx, serviceReq)
	if err != nil {
		return nil, err
	}
	if err := gconv.Struct(user, &res.User); err != nil {
		return nil, err
	}
	if session := service.Context.Get(ctx).Session; session != nil {
		res.Token = session.Id()
	}
	return &res, err
}

// 获取Session信息
func (s *userApi) GetSession(ctx context.Context, req *user.GetSessionReq) (*user.GetSessionRes, error) {
	session := service.Session.GetByToken(ctx, req.Token)
	if session == nil {
		return nil, gerror.New("Session不存在")
	}
	var res user.GetSessionRes
	if err := gconv.Struct(session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// 检测账号唯一性
func (s *userApi) CheckPassport(ctx context.Context, req *user.CheckPassportReq) (*user.CheckPassportRes, error) {
	res := user.CheckPassportRes{}
	res.Ok = service.User.CheckPassport(ctx, req.Passport)
	return &res, nil
}

// 检测昵称唯一性
func (s *userApi) CheckNickName(ctx context.Context, req *user.CheckNickNameReq) (*user.CheckNickNameRes, error) {
	res := user.CheckNickNameRes{}
	res.Ok = service.User.CheckPassport(ctx, req.Nickname)
	return &res, nil
}

// 查询用户信息
func (s *userApi) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserRes, error) {
	res := user.GetUserRes{}
	user, err := service.User.GetUser(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	if err := gconv.Struct(user, &res.User); err != nil {
		return nil, err
	}
	return &res, nil
}
