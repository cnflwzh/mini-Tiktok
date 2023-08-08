package repository

import (
	"mini-Tiktok/biz/entity"
	"mini-Tiktok/config"
)

// 新增评论
func AddComment(userID int64, videoID int64, content string) (int64, string, error) {
	//要返回commentID和createdAt和error
	var comment entity.CommentTable
	comment.UserId = userID
	comment.VideoId = videoID
	comment.Content = content
	err := config.DB.Create(&comment).Error
	return int64(comment.ID), comment.CreatedAt.Format("2006-01-02 15:04:05"), err
}

// 删除评论
func DeleteComment(commentID int64) (int64, string, error) {
	var comment entity.CommentTable
	comment.ID = uint(commentID)
	err := config.DB.Delete(&comment).Error
	return int64(comment.ID), comment.CreatedAt.Format("2006-01-02 15:04:05"), err
}

// 获取视频的评论列表
func GetCommentList(videoID int64) ([]*entity.CommentTable, error) {
	var commentList []*entity.CommentTable
	err := config.DB.Where("video_id = ?", videoID).Find(&commentList).Error
	return commentList, err
}

// 判断当前用户是否是评论的作者
func IsCommentAuthor(userID int64, commentID int64) (bool, error) {
	var comment entity.CommentTable
	err := config.DB.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return false, err
	}
	return comment.UserId == userID, nil
}
