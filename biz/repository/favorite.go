package repository

import (
	"mini-Tiktok/biz/entity"
	"mini-Tiktok/config"
)

// 判断是否已经点赞
func IsFavorite(userId int64, videoId int64) (bool, error) {
	var count int64
	err := config.DB.Model(&entity.Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Count(&count).Error
	return count > 0, err
}

// 删除点赞
func DeleteFavorite(userId int64, videoId int64) error {
	err := config.DB.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&entity.Favorite{}).Error
	if err != nil {
		return err
	}
	// 视频点赞数-1
	//err = config.DB.Model(&entity.Video{}).Where("id = ?", videoId).Update("favorite_count", config.DB.Raw("favorite_count - ?", 1)).Error
	//if err != nil {
	//	return err
	//}
	// 用户点赞数-1
	//err = config.DB.Model(&entity.User{}).Where("id = ?", userId).Update("favorite_count", config.DB.Raw("favorite_count - ?", 1)).Error
	//if err != nil {
	//	return err
	//}
	// 被点赞的视频作者的点赞数-1
	var video entity.Video
	err = config.DB.Where("id = ?", videoId).First(&video).Error
	if err != nil {
		return err
	}
	err = config.DB.Model(&entity.User{}).Where("id = ?", video.UserId).Update("total_favorited", config.DB.Raw("total_favorited - ?", 1)).Error
	return err
}

// 添加点赞
func AddFavorite(userId int64, videoId int64) error {
	err := config.DB.Create(&entity.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}).Error
	if err != nil {
		return err
	}
	// 视频点赞数+1
	//err = config.DB.Model(&entity.Video{}).Where("id = ?", videoId).Update("favorite_count", config.DB.Raw("favorite_count + ?", 1)).Error
	//if err != nil {
	//	return err
	//}
	// 用户点赞数+1
	//err = config.DB.Model(&entity.User{}).Where("id = ?", userId).Update("favorite_count", config.DB.Raw("favorite_count + ?", 1)).Error
	//if err != nil {
	//	return err
	//}
	// 被点赞的视频作者的点赞数+1
	var video entity.Video
	err = config.DB.Where("id = ?", videoId).First(&video).Error
	if err != nil {
		return err
	}
	err = config.DB.Model(&entity.User{}).Where("id = ?", video.UserId).Update("total_favorited", config.DB.Raw("total_favorited + ?", 1)).Error
	return err
}

// 获取点赞列表
func GetFavoriteList(userId int64) ([]entity.Favorite, error) {
	var favorites []entity.Favorite
	err := config.DB.Where("user_id = ?", userId).Find(&favorites).Error
	return favorites, err
}
