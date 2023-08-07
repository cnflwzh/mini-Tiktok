package entity

type UserFollow struct {
	UserID   int64 `gorm:"column:user_id"`
	FollowID int64 `gorm:"column:follow_id"`
}

func (UserFollow) TableName() string {
	return "user_follow"
}
