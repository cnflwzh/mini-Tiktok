// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/protobuf/proto"
	"mini-Tiktok/biz/middleware/jwt"
	"mini-Tiktok/biz/model/user"
	"mini-Tiktok/biz/repository"
	"mini-Tiktok/biz/utils"
)

// GetUser .
// @router /douyin/user [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrorResponse(c, 12001, err.Error())
		return
	}

	resp := new(user.DouyinUserResponse)
	userById, err := repository.GetUserById(req.GetUserId())
	if err != nil {
		utils.SendErrorResponse(c, 12002, err.Error())
		return
	}
	isFollowing, err := repository.IsFollowing(req.GetTokenUserId(), req.GetUserId())
	if err != nil {
		utils.SendErrorResponse(c, 12003, err.Error())
		return
	}
	commonUser := userById.ToCommonUser(isFollowing)

	resp.User = commonUser
	resp.StatusCode = proto.Int32(0)
	resp.StatusMsg = proto.String("用户查找成功")

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Info("test:", string(c.Request.QueryString()))

	resp := &user.DouyinUserRegisterResponse{
		StatusCode: proto.Int32(0),
		StatusMsg:  proto.String(""),
		UserId:     proto.Int64(0),
		Token:      proto.String(""),
	}
	username := req.GetUsername()
	password := req.GetPassword()
	//判断用户名和密码是否为空
	if username == "" || password == "" {
		utils.SendErrorResponse(c, 10001, "用户名或密码为空")
		return
	}
	//去除用户名和密码的空格
	username = utils.TrimSpace(username)
	password = utils.TrimSpace(password)
	//查询用户密码
	credential, userId, err := repository.GetUserCredential(username)
	if credential == "" {
		utils.SendErrorResponse(c, 10002, "用户不存在")
		return
	}
	if err != nil {
		utils.SendErrorResponse(c, 10003, err.Error())
		return
	}
	//判断密码是否正确
	hash := utils.CheckPasswordHash(password, credential)
	if !hash {
		utils.SendErrorResponse(c, 10004, "密码错误")
		return
	}
	//生成token
	token, err := jwt.GenerateToken(username, userId)
	if err != nil {
		utils.SendErrorResponse(c, 10005, err.Error())
		return
	}

	*resp.StatusCode = 0
	*resp.StatusMsg = "登录成功"
	*resp.UserId = userId
	*resp.Token = token

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrorResponse(c, consts.StatusBadRequest, err.Error())
		return
	}
	resp := &user.DouyinUserRegisterResponse{
		StatusCode: proto.Int32(0),
		StatusMsg:  proto.String(""),
		UserId:     proto.Int64(0),
		Token:      proto.String(""),
	}

	username := *req.Username
	password := *req.Password
	//判断用户名和密码是否为空
	if username == "" || password == "" {
		utils.SendErrorResponse(c, 11001, "用户名或密码为空")
		return
	}
	//去除用户名和密码的空格
	username = utils.TrimSpace(username)
	password = utils.TrimSpace(password)
	//判断用户名和密码是否包含非法字符
	if !utils.CheckUsername(username) || !utils.CheckPassword(password) {
		utils.SendErrorResponse(c, 11007, "用户名或密码不满足要求")
		return
	}
	//判断用户名是否已经存在
	exist := repository.CheckUserExist(username)
	if exist {
		utils.SendErrorResponse(c, 11002, "用户名已存在")
		return
	}
	//获取密码加盐哈希
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		utils.SendErrorResponse(c, 11003, err.Error())
		return
	}
	//插入数据库
	//1、在用户信息表中插入
	userId, err := repository.AddUser(username)
	if err != nil {
		utils.SendErrorResponse(c, 11004, err.Error())
		return
	}
	//2、在用户凭证表中插入
	err = repository.AddUserCredential(username, passwordHash, userId)
	if err != nil {
		utils.SendErrorResponse(c, 11005, err.Error())
		return
	}
	token, err := jwt.GenerateToken(username, userId)
	if err != nil {
		utils.SendErrorResponse(c, 11006, err.Error())
		return
	}

	*resp.StatusCode = 0
	*resp.StatusMsg = "注册成功"
	*resp.UserId = userId
	*resp.Token = token
	c.JSON(consts.StatusOK, resp)
}
