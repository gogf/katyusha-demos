package service

import (
	"context"
	"github.com/gogf/gf/net/gtrace"
	"github.com/gogf/gf/os/gsession"
	"github.com/gogf/gf/util/guid"
	"github.com/gogf/katyusha-demos/app/service/user/internal/model"
	"time"
)

// Session管理服务
var Session = &serviceSession{
	ctxKeyGrpc: "token", // 用于Session唯一性识别
}

type serviceSession struct {
	ctxKeyGrpc string
}

const (
	sessionKey = "SessionKey" // 用于Session中的数据存储Key
)

var (
	// 用于Session存储的Storage，测试场景下使用内存，实际业务中往往使用redis
	sessionStorage = gsession.New(24*time.Hour, gsession.NewStorageMemory())
)

// 获取/创建用户Session.
func (s *serviceSession) New(sessionId string) *gsession.Session {
	return sessionStorage.New(sessionId)
}

// 从Context读取SessionToken。
func (s *serviceSession) Token(ctx context.Context) string {
	return gtrace.GetBaggageVar(ctx, s.ctxKeyGrpc).String()
}

// 设置用户Session.
func (s *serviceSession) Set(ctx context.Context, data *model.Session) error {
	customCtx := Context.Get(ctx)
	if customCtx.Session == nil {
		customCtx.Session = s.New(guid.S())
	}
	return customCtx.Session.Set(sessionKey, data)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *serviceSession) GetByToken(ctx context.Context, token string) *model.Session {
	if v := s.New(token).GetVar(sessionKey); !v.IsNil() {
		var session *model.Session
		_ = v.Struct(&session)
		return session
	}
	return nil
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *serviceSession) Get(ctx context.Context) *model.Session {
	customCtx := Context.Get(ctx)
	if customCtx.Session != nil {
		if v := customCtx.Session.GetVar(sessionKey); !v.IsNil() {
			var session *model.Session
			_ = v.Struct(&session)
			return session
		}
	}
	return nil
}

// 删除用户Session（自动读取InComingMap中的token）。
func (s *serviceSession) Remove(ctx context.Context) error {
	customCtx := Context.Get(ctx)
	if customCtx.Session != nil {
		return customCtx.Session.Remove(sessionKey)
	}
	return nil
}

// 通过给定token删除用户Session信息。
func (s *serviceSession) RemoveByToken(ctx context.Context, token string) error {
	return s.New(token).Remove(sessionKey)
}
