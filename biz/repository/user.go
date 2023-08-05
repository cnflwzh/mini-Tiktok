package repository

import (
	"mini-Tiktok/biz/entity"
)

// 通过用户ID获取用户信息
func GetUserById(id int64) (*entity.User, error) {
	var user entity.User
	err := DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func AddUser(username string) (id int64, err error) {
	user := entity.User{
		Name:            username,
		Avater:          "https://p26-passport.byteacctimg.com/img/user-avatar/3ce68fa625dbd76412bfa45fbb454ec0~180x180.awebp?",
		Signature:       "这个人很懒，什么都没有留下",
		BackgroundImage: "https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/bf772c78531a489d8abe3ebbcbb83e70~tplv-k3u1fbpfcp-no-mark:240:240:240:160.awebp?",
	}
	err = DB.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func UpdateUserFavoriteCount(userId int64, favoriteCount int64) error {
	var user entity.User
	err := DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return err
	}
	user.FavoriteCount += favoriteCount
	err = DB.Save(&user).Error
	return err
}
