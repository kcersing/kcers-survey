package errno

import (
	"errors"
	"fmt"
	"kcers-survey/idl_gen/model/base"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

// NewErrNo return ErrNo
func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success      = NewErrNo(int64(base.Err_Success), "success")
	Unauthorized = NewErrNo(int64(401), "Unauthorized")
	ServiceErr   = NewErrNo(int64(501), "service error")

	NoRoute            = NewErrNo(int64(base.Err_NoRoute), "no route")
	NoMethod           = NewErrNo(int64(base.Err_NoMethod), "no method")
	BadRequest         = NewErrNo(int64(base.Err_BadRequest), "bad request")
	ParamsErr          = NewErrNo(int64(base.Err_ParamsErr), "params error")
	AuthorizeFail      = NewErrNo(int64(base.Err_AuthorizeFail), "authorize failed")
	RecordNotFound     = NewErrNo(int64(base.Err_RecordNotFound), "record not found")
	RecordAlreadyExist = NewErrNo(int64(base.Err_RecordAlreadyExist), "record already exist")
	DirtyData          = NewErrNo(int64(base.Err_DirtyData), "dirty data")
	DateErr            = NewErrNo(501, "日期类型传值错误")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {

	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
