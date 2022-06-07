package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Git-YLLong/simple-douyin/douyin/api/rpc"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/extra"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	// 校验token
	claims, err := AuthMiddleware(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1005,
			StatusMsg:  "token解析错误",
		})
		return
	}

	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1005,
			StatusMsg:  "videoID error",
		})
		return
	}

	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1005,
			StatusMsg:  "actionType error",
		})
		return
	}

	userIdStr := c.Query("userID")
	var userId int64
	if len(userIdStr) == 0 {
		userId = claims.UID
	} else {
		userId, err = strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1005,
				StatusMsg:  "userID error",
			})
			return
		}
	}

	err = rpc.FavoriteAction(context.Background(), &extra.DouyinFavoriteActionRequest{
		VideoId:    videoId,
		UserId:     userId,
		ActionType: int32(actionType),
		Token:      token,
	})

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "操作失败",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})

}
