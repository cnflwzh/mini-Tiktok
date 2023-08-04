package entity

import (
	"gorm.io/gorm"
	"time"
)

type Credentials struct {
	UserId    int64  `gorm:"column:user_id;type:bigint(20);not null" json:"user_id"`
	Username  string `gorm:"primarykey;column:username;type:varchar(16);not null;" json:"username"`
	Password  string `gorm:"column:password;type:char(64);not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Credentials) TableName() string {
	return "user_credentials"
}
