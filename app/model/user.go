// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"github.com/gogf/katyusha-demos/app/model/internal"
)

// User is the golang structure for table user.
type User internal.User

// 注册输入参数
type ServiceUserSignUpReq struct {
	Passport string
	Password string
	Nickname string
}
