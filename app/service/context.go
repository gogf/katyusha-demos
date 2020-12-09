package service

import (
	"context"
	"github.com/gogf/katyusha-demos/app/model"
	"github.com/gogf/katyusha/krpc"
)

// 上下文管理服务
var Context = new(serviceContext)

type serviceContext struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *serviceContext) Init(ctx context.Context, customCtx *model.Context) context.Context {
	return context.WithValue(ctx, model.ContextKey, customCtx)
}

// 获取服务间传递的信息，可能为空
func (s *serviceContext) GetContextGrpc(ctx context.Context) *model.ContextGrpc {
	v := krpc.Ctx.IngoingMap(ctx).GetVar(model.ContextKeyGrpc)
	if v.IsNil() {
		return nil
	}
	var contextGrpc *model.ContextGrpc
	_ = v.Struct(&contextGrpc)
	return contextGrpc
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *serviceContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *serviceContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
