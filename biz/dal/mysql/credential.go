package mysql

import (
	"mini-Tiktok/biz/model/common/credential"
)

func CheckUserExist(username string) bool {
	var userCredential credential.Credentials
	err := DB.Where("username = ?", username).First(&userCredential).Error
	if err != nil {
		return false
	}
	return true
}

func AddUserCredential(username string, password string, userId int) error {
	userCredential := credential.Credentials{
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

func GetUserCredential(username string) (string, int, error) {
	var userCredential credential.Credentials
	err := DB.Where("username = ?", username).First(&userCredential).Error
	if err != nil {
		return "", -1, err
	}
	return userCredential.Password, userCredential.UserId, nil
}
