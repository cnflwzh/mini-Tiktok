package entity

import "gorm.io/gorm"

// DROP TABLE IF EXISTS `video_comment`;
// CREATE TABLE `video_comment`  (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
//   `video_id` bigint(20) NOT NULL COMMENT '视频id',
//   `user_id` bigint(20) NOT NULL COMMENT '用户id',
//   `content` varchar(1500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
//   `created_at` datetime NOT NULL COMMENT '创建时间',
//   `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
//   `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
//   PRIMARY KEY (`id`) USING BTREE,
//   INDEX `index_comment_video_id`(`video_id`) USING BTREE,
//   INDEX `fk_comment_user_id`(`user_id`) USING BTREE,
//   CONSTRAINT `fk_comment_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_profile` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
//   CONSTRAINT `fk_comment_video_id` FOREIGN KEY (`video_id`) REFERENCES `video_info` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
// ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '视频评论表' ROW_FORMAT = Dynamic;

type CommentTable struct {
	gorm.Model
	VideoId int64  `gorm:"column:video_id;not null;comment:'视频id'" json:"video_id"`
	UserId  int64  `gorm:"column:user_id;not null;comment:'用户id'" json:"user_id"`
	Content string `gorm:"column:content;not null;type:varchar(1500);comment:'评论内容'" json:"content"`
}

func (CommentTable) TableName() string {
	return "video_comment"
}
