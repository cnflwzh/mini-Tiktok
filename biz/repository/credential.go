package repository

import (
	"mini-Tiktok/biz/entity"
)

func CheckUserExist(username string) bool {
	var userCredential entity.Credentials
	err := DB.Where("username = ?", username).First(&userCredential).Error
	return err == nil
}

func AddUserCredential(username string, password string, userId int64) error {
	userCredential := entity.Credentials{
		Username: username,
		Password: password,
		UserId:   userId,
	}
	err := DB.Create(&userCredential).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserCredential(username string) (string, int64, error) {
	var userCredential entity.Credentials
	err := DB.Where("username = ?", username).First(&userCredential).Error
	if err != nil {
		return "", -1, err
	}
	return userCredential.Password, userCredential.UserId, nil
}
