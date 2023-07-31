package mysql

import (
	"mini-Tiktok/biz/model/common/video"
)

// 获取视频
func GetVideo(videoId int64) (*video.Video, error) {
	var v video.Video
	err := DB.Where("id = ?", videoId).First(&v).Error
	return &v, err
}
