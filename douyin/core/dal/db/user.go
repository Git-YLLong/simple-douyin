package db

import (
	"context"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}

type UserInfo struct {
	//	gorm.Model
	Id            int64  `json:"id"` // ID唯一
	Name          string `json:"name"`
	FollowCount   string `json:"follow_count"`
	FollowerCount string `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// 新增用户
func CreateUser(ctx context.Context, user []*User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// 查询用户
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
