package main

import (
	"context"

	"github.com/Git-YLLong/simple-douyin/douyin/core/pack"
	"github.com/Git-YLLong/simple-douyin/douyin/core/service"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
	"github.com/Git-YLLong/simple-douyin/pkg/jwt"
)

// CoreServiceImpl implements the last service interface defined in the IDL.
type CoreServiceImpl struct{}

// Register implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) Register(ctx context.Context, req *core.DouyinUserRegisterRequest) (resp *core.DouyinUserRegisterResponse, err error) {

	resp = new(core.DouyinUserRegisterResponse)
	resp.UserId = 0
	resp.Token = ""

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	// 注册用户
	if err := pack.Init("2022-05-25", 1); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	var userid = pack.GenId() // 使用雪花算法生成唯一id
	err = service.NewUserService(ctx).Register(req, userid)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	token, err := jwt.CreateJwtToken(userid)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	//err = service.NewUserService(ctx).NewUserInfo()   // 预留一下新增用户信息接口
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = userid
	resp.Token = token

	return resp, nil
}

// Login implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) Login(ctx context.Context, req *core.DouyinUserLoginRequest) (resp *core.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) GetUser(ctx context.Context, req *core.DouyinUserRequest) (resp *core.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoFeed implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) GetVideoFeed(ctx context.Context, req *core.DouyinFeedRequest) (resp *core.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishVideo implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) PublishVideo(ctx context.Context, req *core.DouyinPublishActionRequest) (resp *core.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPublishedList implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) GetPublishedList(ctx context.Context, req *core.DouyinPublishListRequest) (resp *core.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}
