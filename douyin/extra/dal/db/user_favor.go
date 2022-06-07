package db

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserFavor struct {
	VideoID int64 `json:"video_id" gorm:"primaryKey"`
	UserID  int64 `json:"user_id" gorm:"not null"`
	Status  int32 `json:"status" gorm:"default 2"` // 创建默认为1-点赞，取消点赞时为2
}

func QueryInfo(ctx context.Context, userFavor *UserFavor) error {
	return DB.WithContext(ctx).Where("video_id = ? and user_id = ?", userFavor.VideoID, userFavor.UserID).
		Take(&UserFavor{}).Error
}

func AddUserFavor(ctx context.Context, userFavor *UserFavor) error {
	return DB.WithContext(ctx).Create(userFavor).Error
}

func UpdateUserFavor(ctx context.Context, userFavor *UserFavor) error {
	return DB.WithContext(ctx).Model(&UserFavor{}).Where("video_id = ? and user_id = ?", userFavor.VideoID, userFavor.UserID).
		Update("status", userFavor.Status).Error
}

func UpdateVideoFC(ctx context.Context, videoID int64, actionType int32) error {
	if actionType == 1 {
		return DB.WithContext(ctx).Table("videos").Where("id = ?", videoID).
			Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error
	}
	if actionType == 2 {
		return DB.WithContext(ctx).Table("videos").Where("id = ?", videoID).
			Update("favorite_count", gorm.Expr("favorite_count- ?", 1)).Error
	}
	return errors.New("error")
}
