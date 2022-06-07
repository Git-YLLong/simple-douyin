package handlers

import (
	"context"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Git-YLLong/simple-douyin/douyin/api/rpc"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context) {

	token := c.PostForm("token")
	title := c.PostForm("title")
	file, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1005,
			"status_msg":  "请上传数据",
		})
		return
	}

	// 校验token
	claims, err := AuthMiddleware(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1005,
			StatusMsg:  "token解析错误",
		})
		return
	}
	authorId := claims.UID
	time := time.Now().Unix()

	// 保存到本地
	// 作者id + 时间戳生成唯一文件名，只支持mp4格式
	filename := strconv.FormatInt(authorId, 10) + strconv.FormatInt(time, 10) + ".mp4"

	savePath := filepath.Join(constants.RootVideoPath, filename)

	if err = c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "上传失败",
		})
		return
	}

	err = rpc.PublishVideo(context.Background(), &core.DouyinPublishActionRequest{
		Token:    token,
		Title:    title,
		PlayUrl:  filename,
		AuthorId: authorId,
	})

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "Remote Error",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "发布成功",
	})
}
