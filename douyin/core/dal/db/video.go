package db

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Id            int64  `json:"id"`
	AuthorId      int64  `json:"author_id"` // 唯一标记作者信息UserInfo
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommenCount   int64  `json:"commen_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
