// Code generated by hertz generator. DO NOT EDIT.

package relation

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	relation "mini-Tiktok/biz/handler/relation"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			_relation.POST("/action/", append(_relationactionMw(), relation.RelationAction)...)
			{
				_follow := _relation.Group("/follow", _followMw()...)
				_follow.GET("/list/", append(_relationfollowlistMw(), relation.RelationFollowList)...)
			}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				_follower.GET("/list/", append(_relationfollowerlistMw(), relation.RelationFollowerList)...)
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				_friend.GET("/list/", append(_relationfriendlistMw(), relation.RelationFriendList)...)
			}
		}
	}
}
