package errno

// 定义错误信息

import (
	"errors"
	"fmt"
)

const (
	SuccessCode             = 0
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	UserNotExistErrCode     = 10004
	UserAlreadyExistErrCode = 10005
	TokenErrCode            = 10006
	ParamErrCode2           = 10007
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {

	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success             = NewErrNo(SuccessCode, "Success")
	ServiceErr          = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = NewErrNo(ParamErrCode, "用户名或密码不能为空")
	LoginErr            = NewErrNo(LoginErrCode, "密码错误")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "用户不存在")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "用户已经存在")
	TokenErr            = NewErrNo(TokenErrCode, "token错误")
	ParamErr2           = NewErrNo(ParamErrCode2, "参数错误")
)

// 把error转换为error类型
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
