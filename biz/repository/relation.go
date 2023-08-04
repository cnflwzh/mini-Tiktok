package repository

import (
	"fmt"
	"mini-Tiktok/biz/entity"
	"mini-Tiktok/biz/model/common"
	"mini-Tiktok/biz/model/social/relation"
)

func Follow(userId int64, toUserId int64) error {
	var userObj entity.User
	var toUserObj entity.User
	var followRelation relation.Follow
	result := DB.Model(&entity.User{}).Where("user_id = ?", userId).First(&userObj)
	if result.Error != nil {
		// 查询出错或没有找到符合条件的记录
		fmt.Println("Error:", result.Error)
		return result.Error
	}
	result = DB.Model(&entity.User{}).Where("user_id = ?", toUserId).First(&toUserObj)
	if result.Error != nil {
		// 查询出错或没有找到符合条件的记录
		fmt.Println("Error:", result.Error)
		return result.Error
	}
	result = DB.Model(&relation.Follow{}).Where("user_id = ?", userId).Where("follow_id = ?", toUserId).Find(&followRelation)
	if result.Error != nil {
		DB.Model(&relation.Follow{}).Save(relation.Follow{
			UserId:   userId,
			FollowId: toUserId,
		})
		userObj.FollowCount += 1
		toUserObj.FollowerCount += 1
		DB.Save(&userObj)
		DB.Save(&toUserObj)
		fmt.Println("User follow action is successful.")
	} else {
		DB.Model(&relation.Follow{}).Delete(&followRelation)
		userObj.FollowCount -= 1
		toUserObj.FollowerCount -= 1
		DB.Save(&userObj)
		DB.Save(&toUserObj)
		fmt.Println("User cancel follow action is successful.")
	}
	return nil
}

func GetFollowList(userId int64) ([]*common.User, error) {
	var err error
	var users []*entity.User         // 数据库用户列表
	var requiredUsers []*common.User // 标准用户列表
	var followIds []int64            // 用户关注id列表
	err = DB.Model(&relation.Follow{}).Select("follow_id").Where("user_id = ?", userId).Find(&followIds).Error
	if err != nil {
		return nil, err // 这一块的错误处理可以再看一下，我这边是直接返回了不知道会不会有问题
	}
	err = DB.Model(&entity.User{}).Where("user_id in (?)", followIds).Find(&users).Error
	for i := 0; i < len(users); i++ {
		var isFollow bool
		//判断是否当前用户关注了该用户
		result := DB.Model(&relation.Follow{}).Where("user_id = ?", userId).Where("follow_id = ?", users[i].ID).First(nil)
		if result.Error != nil {
			isFollow = false
		} else {
			isFollow = true
		}
		id := users[i].Model.ID
		totalFav := string(rune(users[i].TotalFavorited))
		requiredUsers[i] = &common.User{
			Id:              &id,
			Name:            &users[i].Name,
			FollowCount:     nil,
			FollowerCount:   &users[i].FollowerCount,
			IsFollow:        &isFollow,
			Avatar:          &users[i].Avater,
			BackgroundImage: &users[i].BackgroundImage,
			Signature:       &users[i].Signature,
			TotalFavorited:  &totalFav,
			WorkCount:       nil,
			FavoriteCount:   nil,
		}
	}

	return requiredUsers, err

}

func GetFollowerList(userId int64) ([]*common.User, error) {
	var err error
	var users []*entity.User         // 数据库用户列表
	var requiredUsers []*common.User // 标准用户列表
	var followerIds []int64          // 用户关注id列表
	err = DB.Model(&relation.Follow{}).Select("user_id").Where("follow_id = ?", userId).Find(&followerIds).Error
	if err != nil {
		return nil, err // 这一块的错误处理也可以再看一下
	}
	err = DB.Model(&entity.User{}).Where("user_id in (?)", followerIds).Find(&users).Error
	for i := 0; i < len(users); i++ {
		var isFollow bool
		//判断是否当前用户关注了该用户
		result := DB.Model(&relation.Follow{}).Where("user_id = ?", userId).Where("follow_id = ?", users[i].ID).First(nil)
		if result.Error != nil {
			isFollow = false
		} else {
			isFollow = true
		}
		id := users[i].Model.ID
		totalFav := string(rune(users[i].TotalFavorited))
		requiredUsers[i] = &common.User{
			Id:              &id,
			Name:            &users[i].Name,
			FollowCount:     nil,
			FollowerCount:   &users[i].FollowerCount,
			IsFollow:        &isFollow,
			Avatar:          &users[i].Avater,
			BackgroundImage: &users[i].BackgroundImage,
			Signature:       &users[i].Signature,
			TotalFavorited:  &totalFav,
			WorkCount:       nil,
			FavoriteCount:   nil,
		}
	}

	return requiredUsers, err

}
