package service

import (
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gvalid"
	"google.golang.org/grpc"
)

// 拦截器。
var Interceptor = new(serviceInterceptor)

type serviceInterceptor struct{}

// 自定义上下文拦截器。
func (s *serviceInterceptor) UnaryCtx(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(ctx, req)
}

// 统一StructTag校验器
func (s *serviceInterceptor) UnaryValidate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 如果没有StructTag，那么校验器什么都不会做
	if err := gvalid.CheckStruct(req, nil); err != nil {
		return nil, gerror.Current(err)
	}
	return handler(ctx, req)
}
