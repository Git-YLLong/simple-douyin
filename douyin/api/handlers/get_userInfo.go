package handlers

import (
	"context"
	"strconv"

	"github.com/Git-YLLong/simple-douyin/douyin/api/rpc"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	token := c.Query("token")
	userIdStr := c.Query("user_id")
	var userId int64
	var user UserInfo

	// token校验
	claims, err := AuthMiddleware(token)
	if err != nil {
		GetUserResponse(c, err, user)
		return
	}

	if len(userIdStr) != 0 { // 获取到了userId
		userId, _ = strconv.ParseInt(userIdStr, 10, 64)
	} else { // 没获取到，用token解析
		userId = claims.UID
	}

	userInfo, err := rpc.GetUser(context.Background(), &core.DouyinUserRequest{
		UserId: userId, // 保证传入server端的userId不为空
		Token:  token,
	})

	user = UserInfo{
		Id:            userInfo.Id,
		Name:          userInfo.Name,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      userInfo.IsFollow,
	}

	if err != nil {
		GetUserResponse(c, errno.ConvertErr(err), user)
		return
	}
	GetUserResponse(c, errno.Success, user)
}
