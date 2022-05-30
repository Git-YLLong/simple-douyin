package handlers

import (
	"net/http"

	"github.com/Git-YLLong/simple-douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
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

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
