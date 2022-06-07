package db

import "context"

type Video struct {
	Id            int64  `json:"id" gorm:"primaryKey"`
	AuthorId      int64  `json:"author_id" gorm:"not null"` // 唯一标记作者信息UserInfo
	PlayUrl       string `json:"play_url" gorm:"not null"`
	CoverUrl      string `json:"cover_url" gorm:"not null"`
	FavoriteCount int64  `json:"favorite_count" gorm:"default:0"`
	CommentCount  int64  `json:"comment_count" gorm:"default:0"`
	IsFavorite    bool   `json:"is_favorite" gorm:"default:false"`
	Title         string `json:"title"`
	UpTime        int64  `json:"up_time"` // 投稿时间(gorm可以自定义创建/更新时间追踪)
	// UserInfo      []*UserInfo `json:"user_info" gorm:"many2many:user_favor;foreignKey:Id;joinForeignKey:userId;joinReferences:favorId;"` // 一个视频被多个用户点赞
}

// 视频投稿
func PublishVideo(ctx context.Context, video []*Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

// 返回按投稿时间倒序列表，限制最多30个
func QueryVideoFeed(ctx context.Context, latestTime int64) ([]*Video, error) {
	res := make([]*Video, 0)
	err := DB.WithContext(ctx).Model(&Video{}).Where("up_time<?", latestTime).Order("up_time DESC").Limit(10).Find(&res).Error

	if err != nil {
		return nil, err
	}

	return res, nil
}

func QueryPublishedVideosById(ctx context.Context, userId int64) ([]*Video, error) {
	res := make([]*Video, 0)
	err := DB.Where("author_id=?", userId).Find(&res).Error

	if err != nil {
		return nil, err
	}

	return res, nil
}
