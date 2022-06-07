package db

import (
	"context"
)

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	UserName string `json:"user_name" gorm:"not null"`
	PassWord string `json:"password" gorm:"not null"`
}

type UserInfo struct {
	Id            int64  `json:"id" gorm:"primaryKey"` // 同User.Id
	Name          string `json:"name" gorm:"omitempty"`
	FollowCount   int64  `json:"follow_count" gorm:"omitempty"`
	FollowerCount int64  `json:"follower_count" gorm:"omitempty"`
	IsFollow      bool   `json:"is_follow" gorm:"omitempty"`
	// 一对一关系
	// User User `json:"user" gorm:"foreignKey:Id`
	// 用户-视频多对多关系
	// FavoriteVideos []*Video `json:"favorite_videos" gorm:"many2many:user_favor;foreignKey:Id;joinForeignKey:favorId;joinReferences:userId;"`
}

// 关联初始化建表：DB.AutoMigrate(&User{}, &UserInfo{})

// 创建用户账户
func CreateUser(ctx context.Context, user []*User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// 创建用户信息
func CreateUserInfo(ctx context.Context, userInfo []*UserInfo) error {
	return DB.WithContext(ctx).Create(userInfo).Error
}

// 查询用户账户信息
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	// 避免拼接sql语句，采用预编译，防止sql注入
	if err := DB.WithContext(ctx).Model(&User{}).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// 根据ID查询用户信息
func QueryUserInfo(ctx context.Context, userId int64) (*UserInfo, error) {
	var res *UserInfo

	if err := DB.WithContext(ctx).Model(&UserInfo{}).Where("id = ?", userId).Take(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// 联表查询
func QueryFavoriteVideosById(ctx context.Context, userId int64) ([]*Video, error) {
	var favoriteVideoList []*Video

	// 根据userId查询user_favors数据库，获取点赞的video_id
	// 查询条件 Where("user_id = ? and status = 1", userId)
	// 根据video_id查询videos数据库

	DB.WithContext(ctx).Table("videos").
		// Select("go_service_info.serviceId as service_id, go_service_info.serviceName as service_name, go_system_info.systemId as system_id, go_system_info.systemName as system_name").
		Joins("left join user_favors on user_favors.video_id = videos.id").
		Where("user_favors.user_id = ? and user_favors.status = 1", userId).
		Scan(&favoriteVideoList)

	return favoriteVideoList, nil
}
