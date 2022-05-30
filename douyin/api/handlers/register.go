package handlers

import (
	"context"

	"github.com/Git-YLLong/simple-douyin/douyin/api/rpc"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerVar UserParam
	// Query Params取参数，demo比较奇葩
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	var token = ""

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		LoginResponse(c, errno.ParamErr, 0, token)
		return
	}

	id, token, err := rpc.Register(context.Background(), &core.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		LoginResponse(c, errno.ConvertErr(err), 0, token)
		return
	}
	LoginResponse(c, errno.Success, id, token)
}
