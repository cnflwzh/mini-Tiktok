package repository

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"mini-Tiktok/biz/entity"
	"mini-Tiktok/biz/model/common"
	"mini-Tiktok/biz/model/social/relation"
	"mini-Tiktok/config"
)

func Follow(userId int64, toUserId int64) error {
	if userId == toUserId {
		return errors.New("不能关注自己")
	}
	var userObj entity.User
	var toUserObj entity.User
	var followRelation entity.UserFollow
	result := config.DB.Model(&entity.User{}).Where("id = ?", userId).First(&userObj)
	if result.Error != nil {
		// 查询出错或没有找到符合条件的记录
		hlog.Error("Error:", result.Error)
		return result.Error
	}
	result = config.DB.Model(&entity.User{}).Where("id = ?", toUserId).First(&toUserObj)
	if result.Error != nil {
		// 查询出错或没有找到符合条件的记录
		hlog.Error("Error:", result.Error)
		return result.Error
	}
	result = config.DB.Model(&entity.UserFollow{}).Where("user_id = ?", userId).Where("follow_id = ?", toUserId).Find(&followRelation)
	if followRelation.UserID == 0 {
		followRelation = entity.UserFollow{
			UserID:   userId,
			FollowID: toUserId,
		}
		config.DB.Model(&entity.UserFollow{}).Create(&followRelation)
		userObj.FollowCount = userObj.FollowCount + 1
		toUserObj.FollowerCount = toUserObj.FollowerCount + 1
		config.DB.Save(&userObj)
		config.DB.Save(&toUserObj)
		hlog.Info("User follow action is successful.")
	} else {
		config.DB.Where("user_id = ?", userId).Where("follow_id = ?", toUserId).Unscoped().Delete(&followRelation)
		if userObj.FollowCount == 0 || toUserObj.FollowerCount == 0 {
			hlog.Info("User cancel follow action is unsuccessful.")
			return errors.New("关注操作异常")
		}
		userObj.FollowCount -= 1
		toUserObj.FollowerCount -= 1
		config.DB.Save(&userObj)
		config.DB.Save(&toUserObj)
		hlog.Info("User cancel follow action is successful.")
	}
	return nil
}

func GetFollowList(userId int64) ([]*common.User, error) {
	var err error
	var users []*entity.User         // 数据库用户列表
	var requiredUsers []*common.User // 标准用户列表
	var followIds []int64            // 用户关注id列表
	err = config.DB.Model(&entity.UserFollow{}).Select("follow_id").Where("user_id = ?", userId).Find(&followIds).Error
	if err != nil {
		hlog.Error("follow error:", err.Error())
		return nil, err // 这一块的错误处理可以再看一下，我这边是直接返回了不知道会不会有问题
	}
	if len(followIds) == 0 {
		hlog.Info("关注列表为空")
		return nil, err
	}
	err = config.DB.Model(&entity.User{}).Where("id in (?)", followIds).Find(&users).Error
	requiredUsers = make([]*common.User, len(users))
	for i := 0; i < len(users); i++ {
		var isFollow bool
		//判断是否当前用户关注了该用户
		result := config.DB.Model(&entity.UserFollow{}).Where("user_id = ?", userId).Where("follow_id = ?", users[i].ID).First(nil)
		if result.Error != nil {
			isFollow = false
		} else {
			isFollow = true
		}
		requiredUsers[i] = users[i].ToCommonUser(isFollow)
	}

	return requiredUsers, err

}

func GetFollowerList(userId int64) ([]*common.User, error) {
	var err error
	var users []*entity.User         // 数据库用户列表
	var requiredUsers []*common.User // 标准用户列表
	var followerIds []int64          // 用户关注id列表
	err = config.DB.Model(&entity.UserFollow{}).Select("user_id").Where("follow_id = ?", userId).Find(&followerIds).Error
	if err != nil {
		return nil, err // 这一块的错误处理也可以再看一下
	}
	if len(followerIds) == 0 {
		hlog.Info("粉丝列表为空")
		return nil, err
	}
	err = config.DB.Model(&entity.User{}).Where("id in (?)", followerIds).Find(&users).Error
	requiredUsers = make([]*common.User, len(users))
	for i := 0; i < len(users); i++ {
		var isFollow bool
		//判断是否当前用户关注了该用户
		result := config.DB.Model(&entity.UserFollow{}).Where("user_id = ?", userId).Where("follow_id = ?", users[i].ID).First(nil)
		if result.Error != nil {
			isFollow = false
		} else {
			isFollow = true
		}
		requiredUsers[i] = users[i].ToCommonUser(isFollow)
	}

	return requiredUsers, err

}

func GetFriendList(userId int64) ([]*relation.FriendUser, error) {
	var err error
	var requiredUsers []*relation.FriendUser // 标准朋友列表
	var followerIds []int64                  // 用户粉丝id列表
	var followIds []int64                    // 用户关注id列表
	var friendIds []int64                    // 朋友id列表
	err = config.DB.Model(&entity.UserFollow{}).Select("follow_id").Where("user_id = ?", userId).Find(&followIds).Error
	err = config.DB.Model(&entity.UserFollow{}).Select("user_id").Where("follow_id = ?", userId).Find(&followerIds).Error
	if len(followerIds) == 0 || len(followIds) == 0 {
		hlog.Info("朋友列表为空")
		return nil, err
	}
	err = config.DB.Model(&entity.UserFollow{}).Select("user_id").Where("user_id in (?)", followIds).Where("follow_id = ?", userId).Find(&friendIds).Error
	for _, id := range followerIds {
		isFollowing, _ := IsFollowing(userId, id)
		if isFollowing {
			friendIds = append(friendIds, id)
		}
	}
	requiredUsers = make([]*relation.FriendUser, len(friendIds))
	hlog.Info(friendIds)
	// 更新所有朋友最新消息
	for i := 0; i < len(friendIds); i++ {
		var MsgType int64
		var latestMessage entity.InterMessage
		var user entity.User
		config.DB.Model(&entity.InterMessage{}).Order("created_at desc").
			Where("from_user_id = ?", userId).Where("to_user_id = ?", friendIds[i]).
			Or("from_user_id = ?", friendIds[i]).Where("to_user_id = ?", userId).
			First(&latestMessage)
		if latestMessage.FromUserID == userId {
			MsgType = 1
			config.DB.Model(&entity.User{}).Where("id = ?", latestMessage.ToUserID).First(&user)
		} else {
			MsgType = 0
			config.DB.Model(&entity.User{}).Where("id = ?", latestMessage.FromUserID).First(&user)
		}

		config.DB.Model(&entity.User{}).Where("id = ?", userId).Where("to_user_id = ?", userId).
			First(&latestMessage)
		isFollowing, _ := IsFollowing(userId, user.ID)
		requiredUsers[i] = &relation.FriendUser{
			Message: &latestMessage.Content,
			MsgType: &MsgType,
			User:    *user.ToCommonUser(isFollowing),
		}
	}

	return requiredUsers, err

}

// IsFollowing 查询用户是否关注某人
func IsFollowing(userId int64, followId int64) (bool, error) {
	var followRelation entity.UserFollow
	result := config.DB.Model(&entity.UserFollow{}).
		Where("user_id = ?", userId).
		Where("follow_id = ?", followId).
		First(&followRelation)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // 没有找到记录，说明用户未关注
		}
		return false, result.Error
	}

	return true, nil // 找到记录，说明用户已关注
}
