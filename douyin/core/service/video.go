package service

import (
	"context"
	"time"

	"github.com/Git-YLLong/simple-douyin/douyin/core/dal/db"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
)

type VideoService struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (s *VideoService) PublishVideo(req *core.DouyinPublishActionRequest, id int64) error {

	var video_id = id
	var coverUrl = "cover.jpg" // 默认都是这个算了

	return db.PublishVideo(s.ctx, []*db.Video{{
		Id:            video_id, // 自动生成一个id
		AuthorId:      req.AuthorId,
		PlayUrl:       req.PlayUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,     // 初始点赞数为0
		CommentCount:  0,     // 初始评论数为0
		IsFavorite:    false, // 默认未点赞
		Title:         req.Title,
		UpTime:        time.Now().Unix(),
	}})
}

func (s *VideoService) Feed(req *core.DouyinFeedRequest) ([]*core.Video, int64, error) {

	// 访问videos数据库，获取video信息
	latestTime := req.LatestTime
	// 返回按投稿时间倒序列表，限制最多30个
	videos, err := db.QueryVideoFeed(s.ctx, latestTime)
	if err != nil || len(videos) == 0 { // 空指针
		return nil, 0, err
	}

	nextTime := time.Now().Unix()
	var videoList []*core.Video
	for _, video := range videos {
		// 对应userId访问userInfo
		userInfo, err := db.QueryUserInfo(s.ctx, video.AuthorId)
		if err != nil || userInfo == nil {
			return nil, 0, err
		}
		// 需要根据用户关注关系表更新IsFollow字段
		authorInfo := &core.User{
			Id:            userInfo.Id,
			Name:          userInfo.Name,
			FollowCount:   &userInfo.FollowCount,
			FollowerCount: &userInfo.FollowerCount,
			IsFollow:      userInfo.IsFollow,
		}
		// 需要根据用户点赞关系表更新IsFavorite字段
		videoList = append(videoList, &core.Video{
			Id:            video.Id,
			Author:        authorInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})

		// 本批返回视频中的最早投稿时间
		if nextTime > video.UpTime {
			nextTime = video.UpTime
		}
	}
	return videoList, nextTime, nil
}

func (s *VideoService) GetPublishedList(req *core.DouyinPublishListRequest) ([]*core.Video, error) {
	// 按登录用户id查找发布列表
	videos, err := db.QueryPublishedVideosById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// 对应userId访问userInfo
	userInfo, err := db.QueryUserInfo(s.ctx, req.UserId)
	if err != nil || userInfo == nil {
		return nil, err
	}
	// IsFollow字段待处理
	authorInfo := &core.User{
		Id:            userInfo.Id,
		Name:          userInfo.Name,
		FollowCount:   &userInfo.FollowCount,
		FollowerCount: &userInfo.FollowerCount,
		IsFollow:      userInfo.IsFollow,
	}

	var videoList []*core.Video
	for _, video := range videos {
		// IsFavorite字段待处理
		videoList = append(videoList, &core.Video{
			Id:            video.Id,
			Author:        authorInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}
	return videoList, nil
}
