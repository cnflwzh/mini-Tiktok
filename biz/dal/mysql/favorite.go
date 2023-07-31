package mysql

import (
	"mini-Tiktok/biz/model/interact/favorite"
)

// 判断是否已经点赞
func IsFavorite(userId int64, videoId int64) (bool, error) {
	var count int64
	err := DB.Model(&favorite.Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Count(&count).Error
	return count > 0, err
}

// 删除点赞
func DeleteFavorite(userId int64, videoId int64) error {
	return DB.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&favorite.Favorite{}).Error
}

// 添加点赞
func AddFavorite(userId int64, videoId int64) error {
	return DB.Create(&favorite.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}).Error
}

// 获取点赞列表
func GetFavoriteList(userId int64) ([]favorite.Favorite, error) {
	var favorites []favorite.Favorite
	err := DB.Where("user_id = ?", userId).Find(&favorites).Error
	return favorites, err
}
