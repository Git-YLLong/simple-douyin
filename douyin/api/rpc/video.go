package rpc

import (
	"context"

	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
)

func Feed(ctx context.Context, req *core.DouyinFeedRequest) ([]*core.Video, int64, error) {
	resp, err := userClient.GetVideoFeed(ctx, req)

	if err != nil {
		return nil, 0, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
	}

	return resp.VideoList, resp.NextTime, nil
}

func PublishVideo(ctx context.Context, req *core.DouyinPublishActionRequest) error {
	resp, err := userClient.PublishVideo(ctx, req)

	if err != nil {
		return err
	}

	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
	}

	return nil
}

func GetPublishedList(ctx context.Context, req *core.DouyinPublishListRequest) ([]*core.Video, error) {
	resp, err := userClient.GetPublishedList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
	}

	return resp.VideoList, nil
}

func GetFavoriteList(ctx context.Context, req *core.DouyinFavoriteListRequest) ([]*core.Video, error) {
	resp, err := userClient.GetFavoriteList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, *resp.BaseResp.StatusMsg)
	}

	return resp.VideoList, nil
}
