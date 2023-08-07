package entity

type Follow struct {
	UserId   int64 `gorm:"primarykey;column:user_id"`
	FollowId int64 `gorm:"primarykey;column:follow_id"`
}

func (table *Follow) TableName() string {
	return "user_follow"
}
