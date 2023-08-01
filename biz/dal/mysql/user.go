package mysql

import (
	"mini-Tiktok/biz/model/common/user"
)

// 通过用户ID获取用户信息
func GetUserById(id int64) (*user.User, error) {
	var user user.User
	err := DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
