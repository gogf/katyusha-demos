package model

import "github.com/gogf/gf/os/gtime"

// Session存储数据结构
type Session struct {
	User       *User       // 用户信息
	LoginTime  *gtime.Time // 登录时间
	LoginAgent string      // 登录终端
}
