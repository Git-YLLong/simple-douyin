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
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		LoginResponse(c, errno.ParamErr, 0, "")
		return
	}

	id, token, err := rpc.Register(context.Background(), &core.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})

	if err != nil {
		LoginResponse(c, errno.ConvertErr(err), 0, "")
		return
	}

	LoginResponse(c, errno.Success, id, token)
}
