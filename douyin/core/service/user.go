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
		return errno.UserAlreadyExistErr
	}
	h := md5.New() // 哈希加密
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
