package favorite

// DROP TABLE IF EXISTS `video_user_like`;
// CREATE TABLE `video_user_like`  (
//   `user_id` bigint(20) NOT NULL COMMENT '用户id',
//   `video_id` bigint(20) NOT NULL COMMENT '视频id',
//   PRIMARY KEY (`user_id`, `video_id`) USING BTREE,
//   INDEX `fk_like_video_id`(`video_id`) USING BTREE,
//   CONSTRAINT `fk_like_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_profile` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
//   CONSTRAINT `fk_like_video_id` FOREIGN KEY (`video_id`) REFERENCES `video_info` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户点赞视频' ROW_FORMAT = Dynamic;

// model
type Favorite struct {
	UserId  int64 `gorm:"primaryKey;column:user_id;type:bigint(20);not null" json:"user_id"`
	VideoId int64 `gorm:"primaryKey;column:video_id;type:bigint(20);not null" json:"video_id"`
}

func (Favorite) TableName() string {
	return "video_user_like"
}
