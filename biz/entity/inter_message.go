package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserMessage struct {
	Id         int64          `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id"`
	FromUserId int64          `gorm:"column:from_user_id;type:bigint(20);not null" json:"from_user_id"`
	ToUserId   int64          `gorm:"column:to_user_id;type:bigint(20);not null" json:"to_user_id"`
	Content    string         `gorm:"column:content;type:varchar(1500);not null" json:"content"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserMessage) TableName() string {
	return "inter_message"
}
