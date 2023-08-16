package repository

import (
	"mini-Tiktok/biz/entity"
	"mini-Tiktok/config"
)

func AddInterMessage(fromUserID, toUserID int64, content string) (id int64, err error) {
	message := entity.InterMessage{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
	}
	err = config.DB.Create(&message).Error
	if err != nil {
		return 0, err
	}
	return message.ID, nil
}

func GetInterMessagesByFromUserID(fromUserID int64) ([]entity.InterMessage, error) {
	var messages []entity.InterMessage
	err := config.DB.Where("from_user_id = ?", fromUserID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetInterMessagesByToUserID(toUserID int64) ([]entity.InterMessage, error) {
	var messages []entity.InterMessage
	err := config.DB.Where("to_user_id = ?", toUserID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetInterMessagesByFromAndToUserID(fromUserID, toUserID int64) ([]entity.InterMessage, error) {
	var messages []entity.InterMessage
	err := config.DB.Where("from_user_id = ? AND to_user_id = ?", fromUserID, toUserID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
