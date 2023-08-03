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

func AddUser(username string) (id uint, err error) {
	user := user.User{
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
