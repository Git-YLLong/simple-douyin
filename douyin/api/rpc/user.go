package rpc

import (
	"context"
	"time"

	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core/coreservice"
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
	"github.com/Git-YLLong/simple-douyin/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient coreservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := coreservice.NewClient(
		constants.CoreServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// Client端，调用远程服务方法

// 注册用户，提供账号密码，获取生成的用户ID和Token
func Register(ctx context.Context, req *core.DouyinUserRegisterRequest) (int64, string, error) {
	resp, err := userClient.Register(ctx, req)
	userId := resp.UserId
	token := resp.Token
	if err != nil {
		return userId, token, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return userId, token, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
	}

	return userId, token, nil
}

// 登录，提供账号密码，获取server校验结果
func Login(ctx context.Context, req *core.DouyinUserLoginRequest) (int64, string, error) {

	resp, err := userClient.Login(ctx, req)
	userId := resp.UserId
	token := resp.Token

	if err != nil {
		return userId, token, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return userId, token, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
	}

	return userId, token, nil
}

// 提供用户Id，获取用户信息
func GetUser(ctx context.Context, req *core.DouyinUserRequest) (*core.User, error) {

	resp, err := userClient.GetUser(ctx, req)
	userInfo := resp.User

	if err != nil {
		return userInfo, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return userInfo, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
	}

	return userInfo, nil
}
