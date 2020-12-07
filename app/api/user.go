package api

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/katyusha-demos/app/model"
	"github.com/gogf/katyusha-demos/app/service"
	"github.com/gogf/katyusha-demos/protocol/pb"
)

// 用户服务
var User = &apiUser{}

type apiUser struct{}

// 注册
func (s *apiUser) SignUp(ctx context.Context, req *pb.SignUpReq) (*pb.NullRes, error) {
	var serviceReq *model.ServiceUserSignUpReq
	if err := gconv.Struct(req, &serviceReq); err != nil {
		return nil, gerror.Current(err)
	}
	err := service.User.SignUp(serviceReq)
	return nil, err
}

// 登录
func (s *apiUser) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInRes, error) {
	panic("implement me")
}

// 检测账号唯一性
func (s *apiUser) CheckPassport(ctx context.Context, req *pb.CheckPassportReq) (*pb.NullRes, error) {
	panic("implement me")
}

// 检测昵称唯一性
func (s *apiUser) CheckNickName(ctx context.Context, req *pb.CheckNickNameReq) (*pb.CheckNickNameRes, error) {
	panic("implement me")
}

// 判断给定用户是否已登录
func (s *apiUser) IsSignedIn(ctx context.Context, req *pb.IsSignedInReq) (*pb.IsSignedInRes, error) {
	panic("implement me")
}

// 查询用户信息
func (s *apiUser) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserRes, error) {
	panic("implement me")
}
