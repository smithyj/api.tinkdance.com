package errorx

import (
	"io"
	"tinkdance/pkg/bizcode"
)

type BizError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewBizError(code int, msg string, data interface{}) *BizError {
	return &BizError{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewBizErrorWithCode(code int) *BizError {
	if code == 0 {
		code = bizcode.Fatal
	}
	msg := bizcode.Msg(code)
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

func NewBizErrorWithMsg(msg string) *BizError {
	code := bizcode.Fatal
	if msg == "" || msg == io.EOF.Error() {
		msg = bizcode.Msg(code)
	}
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

func NewBizErrorWithData(data interface{}) *BizError {
	code := bizcode.Fatal
	msg := bizcode.Msg(code)
	return &BizError{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (e *BizError) Error() string {
	return e.Format().Msg
}

func (e *BizError) Format() *BizError {
	if e.Code == 0 {
		// 强制为错误码
		e.Code = bizcode.Fatal
	}
	if e.Msg == "" {
		e.Msg = bizcode.Msg(e.Code)
	}
	return e
}
