package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/protobuf/proto"
	"mini-Tiktok/biz/model/user"
)

// SendErrorResponse 错误响应
func SendErrorResponse(c *app.RequestContext, statusCode int32, message string) {
	resp := &user.DouyinUserRegisterResponse{
		StatusCode: proto.Int32(statusCode),
		StatusMsg:  proto.String(message),
	}
	c.JSON(consts.StatusOK, resp)
}
