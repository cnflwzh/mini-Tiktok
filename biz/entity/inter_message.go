package entity

import (
	"time"

	"mini-Tiktok/biz/model/social/chat"

	"google.golang.org/protobuf/proto"
)

type InterMessage struct {
	ID         int64      `gorm:"column:id;primary_key;autoIncrement;comment:'消息ID'" json:"id"`
	FromUserID int64      `gorm:"column:from_user_id;not null;comment:'发送者ID'" json:"from_user_id"`
	ToUserID   int64      `gorm:"column:to_user_id;not null;comment:'接收者ID'" json:"to_user_id"`
	Content    string     `gorm:"column:content;not null;type:varchar(255);comment:'消息内容'" json:"content"`
	CreatedAt  time.Time  `gorm:"column:created_at;not null;comment:'创建时间'" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;comment:'更新时间'" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;comment:'删除时间'" json:"-"`
}

func (InterMessage) TableName() string {
	return "inter_message"
}

// ConvertInterMessageToMessage converts an InterMessage to a Message.
func ConvertInterMessageToMessage(interMessage *InterMessage) *chat.Message {
	return &chat.Message{
		Id:         proto.Int64(interMessage.ID),
		ToUserId:   proto.Int64(interMessage.ToUserID),
		FromUserId: proto.Int64(interMessage.FromUserID),
		Content:    proto.String(interMessage.Content),
		CreateTime: proto.String(interMessage.CreatedAt.String()), // Convert time.Time to string
	}
}
