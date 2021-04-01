package api

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/katyusha-demos/app/service/demos-service/internal/model"
	"github.com/gogf/katyusha-demos/app/service/demos-service/internal/service"
	"github.com/gogf/katyusha-demos/app/service/demos-service/protobuf/demos"
)

// 用户服务
var User = &userApi{}

type userApi struct{}

// 注册
func (s *userApi) SignUp(ctx context.Context, req *demos.SignUpReq) (*demos.NullRes, error) {
	var (
		res        demos.NullRes
		serviceReq *model.ServiceUserSignUpReq
	)
	if err := gconv.Struct(req, &serviceReq); err != nil {
		return nil, gerror.Current(err)
	}
	err := service.User.SignUp(ctx, serviceReq)
	return &res, err
}

// 登录
func (s *userApi) SignIn(ctx context.Context, req *demos.SignInReq) (*demos.SignInRes, error) {
	var (
		res        = demos.SignInRes{}
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
func (s *userApi) GetSession(ctx context.Context, req *demos.GetSessionReq) (*demos.GetSessionRes, error) {
	session := service.Session.GetByToken(ctx, req.Token)
	if session == nil {
		return nil, gerror.New("Session不存在")
	}
	var res demos.GetSessionRes
	if err := gconv.Struct(session, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// 检测账号唯一性
func (s *userApi) CheckPassport(ctx context.Context, req *demos.CheckPassportReq) (*demos.CheckPassportRes, error) {
	res := demos.CheckPassportRes{}
	res.Ok = service.User.CheckPassport(ctx, req.Passport)
	return &res, nil
}

// 检测昵称唯一性
func (s *userApi) CheckNickName(ctx context.Context, req *demos.CheckNickNameReq) (*demos.CheckNickNameRes, error) {
	res := demos.CheckNickNameRes{}
	res.Ok = service.User.CheckPassport(ctx, req.Nickname)
	return &res, nil
}

// 查询用户信息
func (s *userApi) GetUser(ctx context.Context, req *demos.GetUserReq) (*demos.GetUserRes, error) {
	res := demos.GetUserRes{}
	user, err := service.User.GetUser(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	if err := gconv.Struct(user, &res.User); err != nil {
		return nil, err
	}
	return &res, nil
}
