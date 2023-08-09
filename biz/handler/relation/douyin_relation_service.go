// Code generated by hertz generator.

package relation

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"mini-Tiktok/biz/model/common"
	"mini-Tiktok/biz/repository"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"mini-Tiktok/biz/model/social/relation"
)

// RelationAction .
// @router /douyin/relation/action [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var userId int
	var toUserId int
	id := c.FormValue("user_id")      // 登录用户
	toId := c.FormValue("to_user_id") // 关注用户
	userId, err = strconv.Atoi(string(id))
	toUserId, err = strconv.Atoi(string(toId))
	if err != nil {
		hlog.Error("follow action error:", err.Error())
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.DouyinRelationActionResponse)
	var StatusCode int32
	var StatusMsg string
	err = repository.Follow(int64(userId), int64(toUserId))
	if err != nil {
		StatusCode = -1
		StatusMsg = "用户关注失败" + err.Error()
		SendErrorResponse(c, StatusCode, StatusMsg)
		hlog.Error("follow action error:", err.Error())
		return
	}
	StatusCode = 0
	StatusMsg = "用户" + strconv.FormatInt(int64(userId), 10) + "关注或取关用户" + strconv.FormatInt(int64(toUserId), 10) + "成功"

	resp = &relation.DouyinRelationActionResponse{
		StatusCode: &StatusCode,
		StatusMsg:  &StatusMsg,
	}

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowList .
// @router /douyin/relation/follow/list [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var userId int
	id := c.FormValue("token_user_id")
	userId, err = strconv.Atoi(string(id))
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		hlog.Error("follow error:", err.Error(), string(c.URI().QueryString()))
		return
	}

	resp := new(relation.DouyinRelationFollowListResponse)
	var StatusCode int32
	var StatusMsg string
	var UserList []*common.User

	UserList, err = repository.GetFollowList(int64(userId))
	if err != nil {
		StatusCode = -1
		StatusMsg = "用户" + strconv.FormatInt(int64(userId), 10) + "浏览关注列表失败"
		SendErrorResponse(c, StatusCode, StatusMsg)
		hlog.Error("follow error:", err.Error())
		return
	}
	StatusCode = 0
	StatusMsg = "用户" + strconv.FormatInt(int64(userId), 10) + "正在浏览关注列表"

	resp = &relation.DouyinRelationFollowListResponse{
		StatusCode: &StatusCode,
		StatusMsg:  &StatusMsg,
		UserList:   UserList,
	}

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowerList .
// @router /douyin/relation/follower/list [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var userId int
	id := c.FormValue("token_user_id")
	userId, err = strconv.Atoi(string(id))
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		hlog.Error("follow error:", err.Error(), string(c.URI().QueryString()))
		return
	}

	resp := new(relation.DouyinRelationFollowerListResponse)
	var StatusCode int32
	var StatusMsg string
	var UserList []*common.User

	UserList, err = repository.GetFollowerList(int64(userId))
	if err != nil {
		StatusCode = int32(-1)
		StatusMsg = "用户" + strconv.FormatInt(int64(userId), 10) + "浏览粉丝列表失败"
		SendErrorResponse(c, StatusCode, StatusMsg)
		hlog.Error("follower error:", err.Error())
		return
	}
	StatusCode = int32(0)
	StatusMsg = "用户" + strconv.FormatInt(int64(userId), 10) + "正在浏览粉丝列表"

	resp = &relation.DouyinRelationFollowerListResponse{
		StatusCode: &StatusCode,
		StatusMsg:  &StatusMsg,
		UserList:   UserList,
	}

	c.JSON(consts.StatusOK, resp)
}

// RelationFriendList .
// @router /douyin/relation/friend/list [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req relation.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(relation.DouyinRelationFriendListResponse)

	c.JSON(consts.StatusOK, resp)
}

func SendErrorResponse(c *app.RequestContext, statusCode int32, message string) {
	resp := &relation.DouyinRelationFollowListResponse{
		StatusCode: &statusCode,
		StatusMsg:  &message,
	}
	c.JSON(consts.StatusOK, resp)
}
