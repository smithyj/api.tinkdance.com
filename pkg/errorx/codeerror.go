package errorx

import (
	"fmt"
	"io"
	"tinkdance/pkg/codex"
)

type Error struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func New(code int, msg string, data interface{}) error {
	return &Error{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func WithCode(code int) error {
	if code == 0 {
		code = codex.Error
	}
	msg := codex.Msg(code)
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func WithMsg(msg string) error {
	code := codex.Error
	if msg == "" || msg == io.EOF.Error() {
		msg = codex.Msg(code)
	}
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func WithData(data interface{}) error {
	code := codex.Error
	msg := codex.Msg(code)
	return &Error{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error codex: %v, error msg: %s, error data: %v", e.Code, e.Msg, e.Data)
}

func (e *Error) Response() *Error {
	if e.Code == 0 {
		// 强制为错误码
		e.Code = codex.Error
	}
	if e.Msg == "" {
		e.Msg = codex.Msg(e.Code)
	}
	return e
}
