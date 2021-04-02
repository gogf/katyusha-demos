package model

// 注册请求参数，用于前后端交互参数格式约定
type UserApiSignUpReq struct {
	Passport  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
}

// 登录请求参数，用于前后端交互参数格式约定
type UserApiSignInReq struct {
	Passport string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

// 账号唯一性检测请求参数，用于前后端交互参数格式约定
type UserApiCheckPassportReq struct {
	Passport string `v:"required#账号不能为空"`
}

// 昵称唯一性检测请求参数，用于前后端交互参数格式约定
type UserApiCheckNickNameReq struct {
	Nickname string `v:"required#昵称不能为空"`
}

// 注册输入参数
type UserServiceSignUpReq struct {
	Passport string
	Password string
	Nickname string
}

// 用户登录参数
type UserServiceSignInReq struct {
	Passport string
	Password string
	Agent    string
}

// 用户登录返回参数
type UserServiceSignInRes struct {
	Token string
}
