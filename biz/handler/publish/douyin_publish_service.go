// Code generated by hertz generator.

package publish

import (
	"context"
	"mini-Tiktok/biz/entity"
	"mini-Tiktok/biz/middleware/jwt"
	"mini-Tiktok/biz/model/common"
	"mini-Tiktok/biz/model/publish"
	"mini-Tiktok/biz/repository"
	"mini-Tiktok/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/protobuf/proto"
)

// PublishAction .
// @router /douyin/publish/action [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	formFile, err := c.FormFile("data")
	token := c.FormValue("token")
	title := c.FormValue("title")
	file, err := utils.ReadFile(formFile)

	userId, err := jwt.ParseToken(string(token))
	if err != nil {
		utils.SendErrorResponse(c, 20001, "token 解析失败")
		hlog.Error("token 解析失败", err)
	}
	isMP4Video := utils.IsMp4Video(file)
	if !isMP4Video {
		utils.SendErrorResponse(c, 20002, "上传的视频格式不是 MP4")
		hlog.Error("上传的视频格式不是 MP4", err)
		return
	}
	videoUrl, coverUrl, err := utils.VideoUploadToKodo(file, userId)
	if err != nil {
		utils.SendErrorResponse(c, 20003, "上传视频失败")
		hlog.Error("上传视频失败", err)
		return
	}
	_, err = repository.AddVideo(userId, videoUrl, coverUrl, string(title))
	if err != nil {
		utils.SendErrorResponse(c, 20004, "数据库返回错误")
		hlog.Error("数据库出错", err)
		return
	}

	resp := &publish.DouyinPublishActionResponse{
		StatusCode: proto.Int32(0),
		StatusMsg:  proto.String(""),
	}
	*resp.StatusCode = 0
	*resp.StatusMsg = "发布成功"

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.DouyinPublishListRequest

	// 鉴权用户是否登录
	jwt.CheckLoginMiddleware()(ctx, c)

	// 获取登录用户ID
	loginUserId, ok := c.Get("loginUserId")
	if !ok {
		c.String(consts.StatusInternalServerError, "无法获取登录用户ID")
		return
	}

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	if req.UserId == nil {
		c.String(consts.StatusBadRequest, "查询用户ID不能为空")
		return
	}

	// 从数据库拿用户视频列表
	videoList, err := repository.GetUserVideos(*req.UserId)
	if err != nil {
		c.String(consts.StatusInternalServerError, "获取用户视频失败")
		return
	}

	// 进行类型断言将 loginUserId 转换为 int64
	loginUserIDInt64, ok := loginUserId.(int64)
	if !ok {
		c.String(consts.StatusInternalServerError, "无法获取正确的登录用户ID")
		return
	}

	// 创建响应对象并将获取到的视频列表填充进去
	commonVideoList, err := ConvertVideoListToProto(videoList, loginUserIDInt64)
	if err != nil {
		c.String(consts.StatusInternalServerError, "转换视频数据结构错误")
		return
	}

	resp := &publish.DouyinPublishListResponse{
		StatusCode: proto.Int32(1),
		StatusMsg:  proto.String(""),
		VideoList:  commonVideoList,
	}
	*resp.StatusCode = 0
	*resp.StatusMsg = "获取用户视频列表成功"

	c.JSON(consts.StatusOK, resp)
}

// 将[]*video.Video转换为[]*common.Video
func ConvertVideoListToProto(videoList []*entity.Video, loginUserId int64) ([]*common.Video, error) {
	var commonVideoList []*common.Video

	mysqlAuthor, err := repository.GetUserById(*&videoList[0].UserId)
	if err != nil {
		return nil, err
	}

	author := mysqlAuthor.ToCommonUser(false)

	isFollow, err := repository.IsFollowing(loginUserId, videoList[0].UserId)
	if err != nil {
		return nil, err
	}

	author.IsFollow = &isFollow

	for _, v := range videoList {
		videoId := int64(v.ID)
		isFavorite, err := repository.IsFavorite(loginUserId, videoId)
		if err != nil {
			return commonVideoList, err
		}
		commonVideo := &common.Video{
			Id:            &videoId,
			Author:        author,
			PlayUrl:       &v.PlayUrl,
			CoverUrl:      &v.CoverUrl,
			FavoriteCount: &v.FavoriteCount,
			CommentCount:  &v.CommentCount,
			Title:         &v.Title,
			IsFavorite:    &isFavorite,
		}
		commonVideoList = append(commonVideoList, commonVideo)
	}

	return commonVideoList, nil
}
