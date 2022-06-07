package handlers

import (
	"errors"
	"net/http"

	"github.com/Git-YLLong/simple-douyin/pkg/errno"
	"github.com/Git-YLLong/simple-douyin/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

// 用户注册和登录的返回值
func LoginResponse(c *gin.Context, err error, id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserLoginResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserId:     id,
		Token:      token,
	})
}

type UserInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   *int64 `json:"follow_count,omitempty"`
	FollowerCount *int64 `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow"`
}

type UserInfoResponse struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg,omitempty"`
	User       UserInfo `json:"user"`
}

// 获取用户信息的返回值
func GetUserResponse(c *gin.Context, err error, userInfo UserInfo) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		User:       userInfo,
	})
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64    `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

type FeedResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
	NextTime  int64    `json:"next_time,omitempty"`
}

type PublishedListResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
}

func AuthMiddleware(token string) (*jwt.JwtClaims, error) {
	if token == "" {
		return nil, errors.New("token无效")
	}
	claims, err := jwt.ParseJwtToken(token)
	if err != nil {
		return nil, errors.New("token解析错误")
	}
	return claims, nil
}
