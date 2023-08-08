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
	return config.DB.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&entity.Favorite{}).Error
}

// 添加点赞
func AddFavorite(userId int64, videoId int64) error {
	return config.DB.Create(&entity.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}).Error
}

// 获取点赞列表
func GetFavoriteList(userId int64) ([]entity.Favorite, error) {
	var favorites []entity.Favorite
	err := config.DB.Where("user_id = ?", userId).Find(&favorites).Error
	return favorites, err
}
