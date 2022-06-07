package service

import (
	"context"

	"github.com/Git-YLLong/simple-douyin/douyin/extra/dal/db"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/extra"
)

type UserFavorService struct {
	ctx context.Context
}

// 创建一个NewUser服务
func NewUserFavorService(ctx context.Context) *UserFavorService {
	return &UserFavorService{ctx: ctx}
}

func (s *UserFavorService) AddUserFavor(req *extra.DouyinFavoriteActionRequest) error {
	// 首先查询数据表是否已存在数据
	err := db.QueryInfo(s.ctx, &db.UserFavor{
		VideoID: req.VideoId,
		UserID:  req.UserId,
	})
	if err == nil { // 已有数据，执行更新操作
		// 级联更新videos表FavoriteCount字段
		err := db.UpdateUserFavor(s.ctx, &db.UserFavor{
			VideoID: req.VideoId,
			UserID:  req.UserId,
			Status:  req.ActionType,
		})
		if err != nil {
			return err
		}
		err = db.UpdateVideoFC(s.ctx, req.VideoId, req.ActionType)
		if err != nil {
			return err
		}
		return nil
	} else { // 没有数据，只可能是点赞，并创建
		if req.ActionType == 1 { // 点赞, 级联更新videos表FavoriteCount字段
			err := db.AddUserFavor(s.ctx, &db.UserFavor{
				VideoID: req.VideoId,
				UserID:  req.UserId,
				Status:  req.ActionType,
			})
			if err != nil {
				return err
			}
			err = db.UpdateVideoFC(s.ctx, req.VideoId, req.ActionType)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}
