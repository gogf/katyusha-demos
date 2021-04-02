package service

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gvalid"
	"github.com/gogf/katyusha-demos/app/service/user/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// 拦截器，类似中间件作用。
var Interceptor = new(serviceInterceptor)

type serviceInterceptor struct{}

// 自定义上下文拦截器。
func (s *serviceInterceptor) UnaryCtx(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	var (
		customCtx    model.Context
		contextGrpc  = Context.GetContextGrpc(ctx)
		sessionToken = Session.Token(ctx)
	)
	defer func() {
		if customCtx.Session != nil {
			// 更新Session TTL、同步内容到Session Storage，
			// 以便其他请求获取相同的Session
			customCtx.Session.Close()
		}
	}()
	// 服务间上下文
	if contextGrpc != nil {
		customCtx.Grpc = contextGrpc
		customCtx.User = contextGrpc.User
	}
	// Session.
	if sessionToken != "" {
		customCtx.Session = Session.New(sessionToken)
	}
	// 初始化上下文变量，该变量仅用于当前请求，不跨服务
	ctx = Context.Init(ctx, &customCtx)
	res, err := handler(ctx, req)
	return res, err
}

// 统一StructTag校验器，使用gvalid数据校验组件
func (s *serviceInterceptor) UnaryValidate(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// 如果没有StructTag，那么校验器什么都不会做
	if err := gvalid.CheckStruct(req, nil); err != nil {
		return nil, gerror.NewCode(
			int(codes.InvalidArgument),
			gerror.Current(err).Error(),
		)
	}
	return handler(ctx, req)
}
