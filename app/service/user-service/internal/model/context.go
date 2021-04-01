// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"github.com/gogf/gf/os/gsession"
)

const (
	ContextKey     = "ContextKey"     // 上下文变量存储键名
	ContextKeyGrpc = "ContextKeyGrpc" // 服务间上下文变量存储键名
)

// 内部方法请求上下文结构
type Context struct {
	Session *gsession.Session // 当前Session管理对象
	User    *ContextUser      // 上下文用户信息(与Session中的数据结构不同)
	Grpc    *ContextGrpc      // 服务间上下文信息
}

// 内部方法请求上下文中的用户信息
type ContextUser struct {
	Id       uint   // 用户ID
	Passport string // 用户账号
	Nickname string // 用户名称
}

// 服务间请求上下文中的信息
type ContextGrpc struct {
	User *ContextUser // 服务间上下文用户信息目前是和ContextUser一致，后期可能会不同
}
