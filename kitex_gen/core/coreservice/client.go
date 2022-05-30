// Code generated by Kitex v0.3.1. DO NOT EDIT.

package coreservice

import (
	"context"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, req *core.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *core.DouyinUserRegisterResponse, err error)
	Login(ctx context.Context, req *core.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *core.DouyinUserLoginResponse, err error)
	GetUser(ctx context.Context, req *core.DouyinUserRequest, callOptions ...callopt.Option) (r *core.DouyinUserResponse, err error)
	GetVideoFeed(ctx context.Context, req *core.DouyinFeedRequest, callOptions ...callopt.Option) (r *core.DouyinFeedResponse, err error)
	PublishVideo(ctx context.Context, req *core.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *core.DouyinPublishActionResponse, err error)
	GetPublishedList(ctx context.Context, req *core.DouyinPublishListRequest, callOptions ...callopt.Option) (r *core.DouyinPublishListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCoreServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCoreServiceClient struct {
	*kClient
}

func (p *kCoreServiceClient) Register(ctx context.Context, req *core.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *core.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kCoreServiceClient) Login(ctx context.Context, req *core.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *core.DouyinUserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kCoreServiceClient) GetUser(ctx context.Context, req *core.DouyinUserRequest, callOptions ...callopt.Option) (r *core.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, req)
}

func (p *kCoreServiceClient) GetVideoFeed(ctx context.Context, req *core.DouyinFeedRequest, callOptions ...callopt.Option) (r *core.DouyinFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoFeed(ctx, req)
}

func (p *kCoreServiceClient) PublishVideo(ctx context.Context, req *core.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *core.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishVideo(ctx, req)
}

func (p *kCoreServiceClient) GetPublishedList(ctx context.Context, req *core.DouyinPublishListRequest, callOptions ...callopt.Option) (r *core.DouyinPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishedList(ctx, req)
}
