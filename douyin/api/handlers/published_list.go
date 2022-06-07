package handlers

import (
	"context"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Git-YLLong/simple-douyin/douyin/api/rpc"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"github.com/gin-gonic/gin"
)

func GetPublishedList(c *gin.Context) {
	token := c.Query("token")
	userIdStr := c.Query("user_id")
	var userId int64

	// 校验token
	claims, err := AuthMiddleware(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1005,
			StatusMsg:  "token解析错误",
		})
		return
	}

	if len(userIdStr) != 0 { // 获取到了userId
		userId, _ = strconv.ParseInt(userIdStr, 10, 64)
	} else { // 没获取到，用token解析
		userId = claims.UID
	}

	videoList, err := rpc.GetPublishedList(context.Background(), &core.DouyinPublishListRequest{
		UserId: userId, // 保证传入server端的userId不为空
		Token:  token,
	})

	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 0, StatusMsg: "发布视频列表加载失败"},
		})
	}

	var Videos = make([]*Video, 0)
	for _, video := range videoList {
		// playUrl、coverUrl需要处理一下
		playUrl := filepath.Join(constants.PlayPrefix, video.PlayUrl)
		coverUrl := filepath.Join(constants.PlayPrefix, video.CoverUrl)
		Videos = append(Videos, &Video{
			Id:            video.Id,
			Author:        UserInfo(*video.Author),
			PlayUrl:       playUrl,
			CoverUrl:      coverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}

	c.JSON(http.StatusOK, PublishedListResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "加载发布列表成功"},
		VideoList: Videos,
	})
}
