package main

import (
	"context"

	"github.com/Git-YLLong/simple-douyin/douyin/extra/service"
	"github.com/Git-YLLong/simple-douyin/kitex_gen/extra"
)

// ExtraServiceImpl implements the last service interface defined in the IDL.
type ExtraServiceImpl struct{}

// FavoriteAction implements the ExtraServiceImpl interface.
func (s *ExtraServiceImpl) FavoriteAction(ctx context.Context, req *extra.DouyinFavoriteActionRequest) (resp *extra.DouyinFavoriteActionResponse, err error) {
	resp = new(extra.DouyinFavoriteActionResponse)

	err = service.NewUserFavorService(ctx).AddUserFavor(req)
	if err != nil {
		resp.StatusCode = 10007
		return resp, err
	}
	resp.StatusCode = 0
	return resp, nil
}

// CommentAction implements the ExtraServiceImpl interface.
func (s *ExtraServiceImpl) CommentAction(ctx context.Context, req *extra.DouyinCommentActionRequest) (resp *extra.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the ExtraServiceImpl interface.
func (s *ExtraServiceImpl) CommentList(ctx context.Context, req *extra.DouyinCommentListRequest) (resp *extra.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationAction implements the ExtraServiceImpl interface.
func (s *ExtraServiceImpl) RelationAction(ctx context.Context, req *extra.DouyinRelationActionRequest) (resp *extra.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the ExtraServiceImpl interface.
func (s *ExtraServiceImpl) RelationFollowList(ctx context.Context, req *extra.DouyinRelationFollowListRequest) (resp *extra.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the ExtraServiceImpl interface.
func (s *ExtraServiceImpl) RelationFollowerList(ctx context.Context, req *extra.DouyinRelationFollowerListRequest) (resp *extra.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
