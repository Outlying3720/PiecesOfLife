package logic

import "bookstore/shared"

var (
	errorDuplicateUsername = shared.NewDefaultError("用户名已被注册")
	errorDuplicateMobile   = shared.NewDefaultError("手机号已被注册")

	errorUsernameUnRegister = shared.NewDefaultError("用户名未注册")
	errorIncorrectPassword  = shared.NewDefaultError("用户密码错误")

	errorUserNotFound = shared.NewDefaultError("用户不存在")
)
