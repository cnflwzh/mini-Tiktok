package utils

import (
	"mini-Tiktok/biz/model/common"
	"mini-Tiktok/biz/repository"
	"strconv"
)

// 提取出一个从数据库中获取用户信息并将其转换为common.User的函数
// 这里的用户信息中用户是否关注不进行设置，需要另外获取。
func GetUserInfoFromDb(userID int64) (*common.User, error) {
	// 获取用户信息
	userInfo, err := repository.GetUserById(userID)
	if err != nil {
		return nil, err
	}
	// 设置用户信息
	followCount := userInfo.FollowCount
	workCount := userInfo.WorkCount
	favoriteCount := userInfo.FavoriteCount
	uId := userInfo.ID
	totalFav := strconv.FormatInt(userInfo.TotalFavorited, 10)
	user := common.User{
		Id:              &uId,
		Name:            &userInfo.Name,
		FollowCount:     &followCount,
		FollowerCount:   &userInfo.FollowerCount,
		BackgroundImage: &userInfo.BackgroundImage,
		Signature:       &userInfo.Signature,
		TotalFavorited:  &totalFav,
		WorkCount:       &workCount,
		FavoriteCount:   &favoriteCount,
		Avatar:          &userInfo.Avater,
		IsFollow:        nil, //等待查询用户是否关注的接口
	}
	return &user, nil
}

// 提取出一个从数据库中获取视频信息并将其转换为common.Video的函数,
// 这里的视频信息包括视频作者信息, 其中用户是否关注和视频是否点赞不进
// 行设置，需要另外获取。
func GetVideoInfoFromDb(videoID int64) (*common.Video, error) {
	// 获取视频信息
	videoInfo, err := repository.GetVideo(videoID)
	if err != nil {
		return nil, err
	}
	// 获取视频作者信息
	user, err := GetUserInfoFromDb(videoInfo.UserId)
	if err != nil {
		return nil, err
	}
	// 设置视频信息
	vId := videoInfo.ID
	video := common.Video{
		Id:            &vId,
		Author:        user,
		PlayUrl:       &videoInfo.PlayUrl,
		CoverUrl:      &videoInfo.CoverUrl,
		FavoriteCount: &videoInfo.FavoriteCount,
		CommentCount:  &videoInfo.CommentCount,
		IsFavorite:    nil, // 这里不进行设置
		Title:         &videoInfo.Title,
	}

	return &video, nil
}
