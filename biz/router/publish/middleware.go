// Code generated by hertz generator.

package publish

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

func _publishMw() []app.HandlerFunc {
	// your code...
	return jwt.JWTAuthMiddleware()
}

func _publishactionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistMw() []app.HandlerFunc {
	// your code...
	return nil
}
