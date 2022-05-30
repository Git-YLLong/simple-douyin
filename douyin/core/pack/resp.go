package pack

import (
	"errors"

	"github.com/Git-YLLong/simple-douyin/kitex_gen/core"
	"github.com/Git-YLLong/simple-douyin/pkg/errno"
)

func BuildBaseResp(err error) *core.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *core.BaseResp {
	resp := &core.BaseResp{StatusCode: err.ErrCode}
	if len(err.ErrMsg) != 0 {
		resp.StatusMsg = &err.ErrMsg
	}
	return resp
}
