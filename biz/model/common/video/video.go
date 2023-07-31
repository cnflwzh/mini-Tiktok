package video

import "time"

// DROP TABLE IF EXISTS `video_info`;
// CREATE TABLE `video_info`  (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '视频ID',
//   `user_id` bigint(20) NOT NULL COMMENT '用户ID',
//   `play_url` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频播放地址',
//   `cover_url` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频封面地址',
//   `favorite_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '视频的点赞总数',
//   `comment_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '视频的评论总数',
//   `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频标题',
//   `created_at` datetime NOT NULL COMMENT '创建时间',
//   `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
//   `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
//   PRIMARY KEY (`id`) USING BTREE,
//   INDEX `user_id`(`user_id`) USING BTREE,
//   CONSTRAINT `fk_video_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_profile` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
// ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '视频信息表' ROW_FORMAT = Dynamic;

// model
type Video struct {
	Id            int64     `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id"`
	UserId        int64     `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`
	PlayUrl       string    `gorm:"column:play_url;type:varchar(300);not null" json:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url;type:varchar(300);not null" json:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count;type:bigint(20);not null" json:"favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count;type:bigint(20);not null" json:"comment_count"`
	Title         string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}
