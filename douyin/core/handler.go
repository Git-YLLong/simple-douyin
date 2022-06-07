package main

import (
	"context"

	"github.com/Git-YLLong/simple-douyin/douyin/core/pack"
	"github.com/Git-YLLong/simple-douyin/douyin/core/service"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
	"github.com/Git-YLLong/simple-douyin/pkg/jwt"
)

// CoreServiceImpl implements the last service interface defined in the IDL.
type CoreServiceImpl struct{}

// Register implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) Register(ctx context.Context, req *core.DouyinUserRegisterRequest) (resp *core.DouyinUserRegisterResponse, err error) {

	resp = new(core.DouyinUserRegisterResponse)

	// if len(req.Username) == 0 || len(req.Password) == 0 {
	// 	resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
	// 	return resp, nil
	// }

	// 注册用户
	if err := pack.Init("2022-05-25", 1); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	// 使用雪花算法生成唯一id
	var userid = pack.GenId()
	err = service.NewUserService(ctx).Register(req, userid)
	if err != nil { // 生成id错误
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = userid

	// 生成token
	token, err := jwt.CreateJwtToken(userid)
	if err != nil { // 生成token错误
		resp.BaseResp = pack.BuildBaseResp(errno.TokenErr)
		return resp, nil
	}
	resp.Token = token

	// 注册用户同时新增用户信息
	err = service.NewUserService(ctx).CreateUserInfo(req, userid)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil
}

// Login implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) Login(ctx context.Context, req *core.DouyinUserLoginRequest) (resp *core.DouyinUserLoginResponse, err error) {

	resp = new(core.DouyinUserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.LoginErr)
		return resp, errno.LoginErr
	}

	// 用户登录
	userid, err := service.NewUserService(ctx).Login(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}

	token, err := jwt.CreateJwtToken(userid)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.TokenErr)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = userid
	resp.Token = token

	return resp, nil
}

// GetUser implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) GetUser(ctx context.Context, req *core.DouyinUserRequest) (resp *core.DouyinUserResponse, err error) {

	resp = new(core.DouyinUserResponse)

	userId := req.UserId
	if userId == 0 {

		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	userInfo, err := service.NewUserService(ctx).GetUser(userId)
	if err != nil { // 数据库访问错误或者用户不存在
		resp.BaseResp = pack.BuildBaseResp(errno.UserNotExistErr)
		return resp, err
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = &core.User{
		Id:            userInfo.Id,
		Name:          userInfo.Name,
		FollowCount:   &userInfo.FollowCount,
		FollowerCount: &userInfo.FollowerCount,
		IsFollow:      userInfo.IsFollow,
	}

	return resp, nil
}

// GetVideoFeed implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) GetVideoFeed(ctx context.Context, req *core.DouyinFeedRequest) (resp *core.DouyinFeedResponse, err error) {
	resp = new(core.DouyinFeedResponse)

	videos, nextTime, err := service.NewVideoService(ctx).Feed(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	if len(videos) == 0 { // 先测试一下能不能取到数据，能得话就删掉
		resp.BaseResp.StatusCode = 1234
		return resp, errno.ServiceErr
	}

	for _, video := range videos {
		resp.VideoList = append(resp.VideoList, &core.Video{
			Id:            video.Id,
			Author:        video.Author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}
	resp.NextTime = nextTime // 选择返回视频里投稿最早的时间戳

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// PublishVideo implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) PublishVideo(ctx context.Context, req *core.DouyinPublishActionRequest) (resp *core.DouyinPublishActionResponse, err error) {

	resp = new(core.DouyinPublishActionResponse)

	// 生成视频id
	if err := pack.Init("2022-05-28", 1); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	var videoId = pack.GenId()

	err = service.NewVideoService(ctx).PublishVideo(req, videoId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetPublishedList implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) GetPublishedList(ctx context.Context, req *core.DouyinPublishListRequest) (resp *core.DouyinPublishListResponse, err error) {
	resp = new(core.DouyinPublishListResponse)

	videos, err := service.NewVideoService(ctx).GetPublishedList(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}

	for _, video := range videos {
		resp.VideoList = append(resp.VideoList, &core.Video{
			Id:            video.Id,
			Author:        video.Author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFavoriteList implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) GetFavoriteList(ctx context.Context, req *core.DouyinFavoriteListRequest) (resp *core.DouyinFavoriteListResponse, err error) {
	resp = new(core.DouyinFavoriteListResponse)

	videos, err := service.NewUserService(ctx).GetFavoriteList(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}

	for _, video := range videos {
		resp.VideoList = append(resp.VideoList, &core.Video{
			Id:            video.Id,
			Author:        video.Author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
