// Code generated by hertz generator.

package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"mini-Tiktok/biz/middleware/jwt"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserMw() []app.HandlerFunc {
	// your code...
	return jwt.JWTAuthMiddleware()
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}
