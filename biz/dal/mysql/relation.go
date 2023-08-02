package mysql

import (
	"mini-Tiktok/biz/model/common"
	"mini-Tiktok/biz/model/common/user"
	"mini-Tiktok/biz/model/social/relation"
)

func GetFollowList(userId int64) ([]*common.User, error) {
	var err error
	var users []*user.User           // 数据库用户列表
	var requiredUsers []*common.User // 标准用户列表
	var followIds []int64            // 用户关注id列表
	err = DB.Model(&relation.Follow{}).Select("follow_id").Where("user_id = ?", userId).Find(&followIds).Error
	err = DB.Model(&user.User{}).Where("user_id in (?)", followIds).Find(&users).Error
	for i := 0; i < len(users); i++ {
		var isFollow bool
		//判断是否当前用户关注了该用户
		result := DB.Model(&relation.Follow{}).Where("user_id = ?", userId).Where("user_id = ?", userId).First(nil)
		if result.Error != nil {
			isFollow = false
		} else {
			isFollow = true
		}
		requiredUsers[i] = &common.User{
			Id:              &users[i].Id,
			Name:            &users[i].Name,
			FollowCount:     nil,
			FollowerCount:   &users[i].FollowerCount,
			IsFollow:        &isFollow,
			Avatar:          &users[i].Avater,
			BackgroundImage: &users[i].BackgroundImage,
			Signature:       &users[i].Signature,
			TotalFavorited:  &users[i].TotalFavorited,
			WorkCount:       nil,
			FavoriteCount:   nil,
		}
	}

	return requiredUsers, err

}

func GetFollowerList(userId int64) ([]*common.User, error) {
	var err error
	var users []*user.User           // 数据库用户列表
	var requiredUsers []*common.User // 标准用户列表
	var followerIds []int64          // 用户关注id列表
	err = DB.Model(&relation.Follow{}).Select("user_id").Where("follow_id = ?", userId).Find(&followerIds).Error
	err = DB.Model(&user.User{}).Where("user_id in (?)", followerIds).Find(&users).Error
	for i := 0; i < len(users); i++ {
		var isFollow bool
		//判断是否当前用户关注了该用户
		result := DB.Model(&relation.Follow{}).Where("user_id = ?", userId).Where("user_id = ?", userId).First(nil)
		if result.Error != nil {
			isFollow = false
		} else {
			isFollow = true
		}
		requiredUsers[i] = &common.User{
			Id:              &users[i].Id,
			Name:            &users[i].Name,
			FollowCount:     nil,
			FollowerCount:   &users[i].FollowerCount,
			IsFollow:        &isFollow,
			Avatar:          &users[i].Avater,
			BackgroundImage: &users[i].BackgroundImage,
			Signature:       &users[i].Signature,
			TotalFavorited:  &users[i].TotalFavorited,
			WorkCount:       nil,
			FavoriteCount:   nil,
		}
	}

	return requiredUsers, err

}
