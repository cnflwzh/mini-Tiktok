package user

import "time"

// DROP TABLE IF EXISTS `user_profile`;
// CREATE TABLE `user_profile`  (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
//   `name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
//   `follow_count` int(11) NOT NULL DEFAULT 0 COMMENT '关注总数',
//   `follower_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '粉丝总数',
//   `avater` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户头像',
//   `background_image` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户个人页顶部大图',
//   `signature` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人简介',
//   `total_favorited` bigint(20) NOT NULL DEFAULT 0 COMMENT '获赞总数',
//   `work_count` int(11) NOT NULL DEFAULT 0 COMMENT '作品数',
//   `favorite_count` int(11) NOT NULL DEFAULT 0 COMMENT '喜欢数',
//   `created_at` datetime NOT NULL COMMENT '创建时间',
//   `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
//   `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
//   PRIMARY KEY (`id`) USING BTREE
// ) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户基本信息表' ROW_FORMAT = Dynamic;

// model
type User struct {
	Id              int64     `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id"`
	Name            string    `gorm:"column:name;type:varchar(16);not null" json:"name"`
	FollowCount     int       `gorm:"column:follow_count;type:int(11);not null" json:"follow_count"`
	FollowerCount   int64     `gorm:"column:follower_count;type:bigint(20);not null" json:"follower_count"`
	Avater          string    `gorm:"column:avater;type:varchar(300)" json:"avater"`
	BackgroundImage string    `gorm:"column:background_image;type:varchar(300)" json:"background_image"`
	Signature       string    `gorm:"column:signature;type:varchar(300)" json:"signature"`
	TotalFavorited  int64     `gorm:"column:total_favorited;type:bigint(20);not null" json:"total_favorited"`
	WorkCount       int       `gorm:"column:work_count;type:int(11);not null" json:"work_count"`
	FavoriteCount   int       `gorm:"column:favorite_count;type:int(11);not null" json:"favorite_count"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}