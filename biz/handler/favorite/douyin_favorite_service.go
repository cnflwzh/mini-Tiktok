// Code generated by hertz generator.

package favorite

import (
	"context"
	dal "mini-Tiktok/biz/repository"

	"mini-Tiktok/biz/utils"

	"mini-Tiktok/biz/model/common"
	"mini-Tiktok/biz/model/interact/favorite"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Action .
// @router /douyin/favorite/action [POST]
func Action(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	//resp := new(favorite.DouyinFavoriteActionResponse)
	if err != nil {
		sendResponse(c, 1, err.Error())
		return
	}
	// 判断当前用户是否已经点赞
	isFavorite, err := dal.IsFavorite(*req.UserId, *req.VideoId)
	if err != nil {
		sendResponse(c, 1, err.Error())
		return
	}
	if err != nil {
		sendResponse(c, 1, err.Error())
		return
	}
	// 判断当前请求是否需要执行
	if *req.ActionType == 1 && isFavorite {
		sendResponse(c, 2, "不能重复点赞")
	} else if *req.ActionType == 2 && !isFavorite {
		sendResponse(c, 3, "不能取消未点赞的视频")
	} else if *req.ActionType == 1 && !isFavorite {
		// 添加点赞
		err = dal.AddFavorite(*req.UserId, *req.VideoId)
		if err != nil {
			sendResponse(c, 4, err.Error())
			return
		}
		sendResponse(c, 0, "点赞成功")

	} else if *req.ActionType == 2 && isFavorite {
		// 取消点赞
		err = dal.DeleteFavorite(*req.UserId, *req.VideoId)
		if err != nil {
			sendResponse(c, 5, err.Error())
			return
		}
		sendResponse(c, 0, "取消点赞成功")
	}
}

// List .
// @router /douyin/favorite/list [GET]
func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	resp := new(favorite.DouyinFavoriteListResponse)
	resp.VideoList = make([]*common.Video, 0)
	if err != nil {
		sendResponse(c, 1, err.Error())
		return
	}
	// 获取点赞列表
	favorites, err := dal.GetFavoriteList(*req.UserId)
	if err != nil {
		sendResponse(c, 1, err.Error())
		return
	}
	// 获取点赞视频列表
	commonVideos := make([]*common.Video, 0)
	for _, favorite := range favorites {
		// 获取视频信息
		videoInfo, err := utils.GetVideoInfoFromDb(favorite.VideoId)
		if err != nil {
			sendResponse(c, 1, err.Error())
			return
		}
		commonVideos = append(commonVideos, videoInfo)
	}
	// 设置响应
	statusCode := int32(0)
	statusMsg := "success"
	resp = &favorite.DouyinFavoriteListResponse{
		StatusCode: &statusCode,
		StatusMsg:  &statusMsg,
		VideoList:  commonVideos,
	}
	c.JSON(consts.StatusOK, resp)
}

func sendResponse(c *app.RequestContext, code int32, msg string) {
	resp := &favorite.DouyinFavoriteActionResponse{
		StatusCode: &code,
		StatusMsg:  &msg,
	}
	c.JSON(consts.StatusOK, resp)
}
