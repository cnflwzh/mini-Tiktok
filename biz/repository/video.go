package repository

import (
	"mini-Tiktok/biz/entity"
)

// 获取视频
func GetVideo(videoId int64) (*entity.Video, error) {
	var v entity.Video
	err := DB.Where("id = ?", videoId).First(&v).Error
	return &v, err
}

func AddVideo(userId int64, videoUrl string, coverUrl string, title string) (int64, error) {
	video := entity.Video{
		UserId:   userId,
		PlayUrl:  videoUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}
	err := DB.Create(&video).Error
	if err != nil {
		return 0, err
	}
	return video.ID, nil
}

func UpdateVideoFavoriteCount(videoId int64, favoriteCount int64) error {
	var video entity.Video
	err := DB.Where("id = ?", videoId).First(&video).Error
	if err != nil {
		return err
	}
	video.FavoriteCount += favoriteCount
	err = DB.Save(&video).Error
	return err
}
