// Code generated by Kitex v0.3.1. DO NOT EDIT.

package coreservice

import (
	"context"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return coreServiceServiceInfo
}

var coreServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CoreService"
	handlerType := (*core.CoreService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":         kitex.NewMethodInfo(registerHandler, newCoreServiceRegisterArgs, newCoreServiceRegisterResult, false),
		"Login":            kitex.NewMethodInfo(loginHandler, newCoreServiceLoginArgs, newCoreServiceLoginResult, false),
		"GetUser":          kitex.NewMethodInfo(getUserHandler, newCoreServiceGetUserArgs, newCoreServiceGetUserResult, false),
		"GetVideoFeed":     kitex.NewMethodInfo(getVideoFeedHandler, newCoreServiceGetVideoFeedArgs, newCoreServiceGetVideoFeedResult, false),
		"PublishVideo":     kitex.NewMethodInfo(publishVideoHandler, newCoreServicePublishVideoArgs, newCoreServicePublishVideoResult, false),
		"GetPublishedList": kitex.NewMethodInfo(getPublishedListHandler, newCoreServiceGetPublishedListArgs, newCoreServiceGetPublishedListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "core",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*core.CoreServiceRegisterArgs)
	realResult := result.(*core.CoreServiceRegisterResult)
	success, err := handler.(core.CoreService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCoreServiceRegisterArgs() interface{} {
	return core.NewCoreServiceRegisterArgs()
}

func newCoreServiceRegisterResult() interface{} {
	return core.NewCoreServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*core.CoreServiceLoginArgs)
	realResult := result.(*core.CoreServiceLoginResult)
	success, err := handler.(core.CoreService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCoreServiceLoginArgs() interface{} {
	return core.NewCoreServiceLoginArgs()
}

func newCoreServiceLoginResult() interface{} {
	return core.NewCoreServiceLoginResult()
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*core.CoreServiceGetUserArgs)
	realResult := result.(*core.CoreServiceGetUserResult)
	success, err := handler.(core.CoreService).GetUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCoreServiceGetUserArgs() interface{} {
	return core.NewCoreServiceGetUserArgs()
}

func newCoreServiceGetUserResult() interface{} {
	return core.NewCoreServiceGetUserResult()
}

func getVideoFeedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*core.CoreServiceGetVideoFeedArgs)
	realResult := result.(*core.CoreServiceGetVideoFeedResult)
	success, err := handler.(core.CoreService).GetVideoFeed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCoreServiceGetVideoFeedArgs() interface{} {
	return core.NewCoreServiceGetVideoFeedArgs()
}

func newCoreServiceGetVideoFeedResult() interface{} {
	return core.NewCoreServiceGetVideoFeedResult()
}

func publishVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*core.CoreServicePublishVideoArgs)
	realResult := result.(*core.CoreServicePublishVideoResult)
	success, err := handler.(core.CoreService).PublishVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCoreServicePublishVideoArgs() interface{} {
	return core.NewCoreServicePublishVideoArgs()
}

func newCoreServicePublishVideoResult() interface{} {
	return core.NewCoreServicePublishVideoResult()
}

func getPublishedListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*core.CoreServiceGetPublishedListArgs)
	realResult := result.(*core.CoreServiceGetPublishedListResult)
	success, err := handler.(core.CoreService).GetPublishedList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCoreServiceGetPublishedListArgs() interface{} {
	return core.NewCoreServiceGetPublishedListArgs()
}

func newCoreServiceGetPublishedListResult() interface{} {
	return core.NewCoreServiceGetPublishedListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *core.DouyinUserRegisterRequest) (r *core.DouyinUserRegisterResponse, err error) {
	var _args core.CoreServiceRegisterArgs
	_args.Req = req
	var _result core.CoreServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *core.DouyinUserLoginRequest) (r *core.DouyinUserLoginResponse, err error) {
	var _args core.CoreServiceLoginArgs
	_args.Req = req
	var _result core.CoreServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, req *core.DouyinUserRequest) (r *core.DouyinUserResponse, err error) {
	var _args core.CoreServiceGetUserArgs
	_args.Req = req
	var _result core.CoreServiceGetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoFeed(ctx context.Context, req *core.DouyinFeedRequest) (r *core.DouyinFeedResponse, err error) {
	var _args core.CoreServiceGetVideoFeedArgs
	_args.Req = req
	var _result core.CoreServiceGetVideoFeedResult
	if err = p.c.Call(ctx, "GetVideoFeed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishVideo(ctx context.Context, req *core.DouyinPublishActionRequest) (r *core.DouyinPublishActionResponse, err error) {
	var _args core.CoreServicePublishVideoArgs
	_args.Req = req
	var _result core.CoreServicePublishVideoResult
	if err = p.c.Call(ctx, "PublishVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPublishedList(ctx context.Context, req *core.DouyinPublishListRequest) (r *core.DouyinPublishListResponse, err error) {
	var _args core.CoreServiceGetPublishedListArgs
	_args.Req = req
	var _result core.CoreServiceGetPublishedListResult
	if err = p.c.Call(ctx, "GetPublishedList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}