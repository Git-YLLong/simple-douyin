package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/Git-YLLong/simple-douyin/douyin/core/dal/db"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
)

type UserService struct {
	ctx context.Context
}

// 创建一个NewUser服务
func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

// 注册用户
func (s *UserService) Register(req *core.DouyinUserRegisterRequest, id int64) error {
	user, err := db.QueryUser(s.ctx, req.Username) // 查询数据库是否已存在用户
	if err != nil {
		return err
	}
	if len(user) != 0 {
		return errno.UserAlreadyExistErr // 用户已存在
	}

	// 哈希加密密码
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}

	passWord := fmt.Sprintf("%x", h.Sum(nil)) // 加密结果转字符串
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.Username,
		PassWord: passWord,
		Id:       id,
	}})
}

// 新增用户信息
func (s *UserService) CreateUserInfo(req *core.DouyinUserRegisterRequest, id int64) error {
	return db.CreateUserInfo(s.ctx, []*db.UserInfo{{
		Id:            id,           // id同账号密码表
		Name:          req.Username, // 昵称默认为用户名
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}})
}

// 用户登录
func (s *UserService) Login(req *core.DouyinUserLoginRequest) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.Username) // 查询数据库是否存在用户
	if err != nil {
		return 0, err
	}
	// 校验用户名
	if len(users) == 0 { // 用户不存在
		return 0, errno.UserNotExistErr
	}

	// 校验密码
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	u := users[0]
	if u.PassWord != passWord { // 密码错误
		return 0, errno.LoginErr
	}
	return u.Id, nil
}

// 获取登陆用户信息
func (s *UserService) GetUser(userId int64) (*db.UserInfo, error) {
	userInfo, err := db.QueryUserInfo(s.ctx, userId)
	if err != nil || userInfo == nil { // 没有用户
		return nil, err
	}
	return userInfo, err
}

func (s *UserService) GetFavoriteList(req *core.DouyinFavoriteListRequest) ([]*core.Video, error) {
	// 访问videos数据库，根据视频id获取video信息
	videos, err := db.QueryFavoriteVideosById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var videoList []*core.Video
	for _, video := range videos {
		// 对应userId访问userInfo
		userInfo, err := db.QueryUserInfo(s.ctx, video.AuthorId)
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
		// IsFavorite字段待处理
		videoList = append(videoList, &core.Video{
			Id:            video.Id,
			Author:        authorInfo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true, // 返回的肯定是点赞的
			Title:         video.Title,
		})
	}
	return videoList, nil
}
