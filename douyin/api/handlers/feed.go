package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Git-YLLong/simple-douyin/douyin/api/rpc"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"github.com/gin-gonic/gin"
)

// 视频流，不限制登录状态
// 刷一个视频调用一次？
func Feed(c *gin.Context) {

	// 返回视频的最新投稿日期，精确到秒
	latestTimeStr := c.Query("latest_time")
	token := c.Query("token")

	var latestTime int64
	// 字符串转int64
	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if err != nil {
		latestTime = time.Now().Unix()
	}
	if latestTime == 0 { // 没有提供latestTime参数，默认为当前时间
		latestTime = time.Now().Unix()
	}

	// 调用Feed服务，返回视频列表，以及下一次获取视频流的latestTime
	videoList, nextTime, err := rpc.Feed(context.Background(), &core.DouyinFeedRequest{
		LatestTime: latestTime,
		Token:      token,
	})

	if err != nil || len(videoList) == 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 0, StatusMsg: "视频返回失败"},
		})
	}

	var Videos = make([]*Video, 0)
	for _, video := range videoList {
		// playUrl、coverUrl需要处理一下
		playUrl := constants.PlayPrefix + video.PlayUrl
		coverUrl := constants.PlayPrefix + video.CoverUrl
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

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "返回成功"},
		VideoList: Videos,
		NextTime:  nextTime, // 是不是空无所谓了，反正下一次调用会处理latestTime
	})
}
