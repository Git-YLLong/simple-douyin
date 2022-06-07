package rpc

import (
	"context"
	"time"

	"github.com/Git-YLLong/simple-douyin/kitex_gen/extra"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/extra/extraservice"
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
	"github.com/Git-YLLong/simple-douyin/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var extraClient extraservice.Client

func initExtraRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := extraservice.NewClient(
		constants.ExtraServiceName,
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
	extraClient = c
}

func FavoriteAction(ctx context.Context, req *extra.DouyinFavoriteActionRequest) error {
	resp, _ := extraClient.FavoriteAction(ctx, req)

	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, *resp.StatusMsg) // 10007
	}
	return nil // 0
}
