package relation

type Follow struct {
	UserId   int64 `gorm:"column:user_id"`
	FollowId int64 `gorm:"column:follow_id"`
}
