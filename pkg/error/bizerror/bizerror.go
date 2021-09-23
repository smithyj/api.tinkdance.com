package bizerror

import (
	"tinkdance/pkg/code/bizcode"
)

type Error struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return e.Format().Msg
}

func (e *Error) Format() *Error {
	if e.Code == 0 {
		// 强制为错误码
		e.Code = bizcode.Fatal
	}
	if e.Msg == "" {
		e.Msg = bizcode.Msg(e.Code)
	}
	return e
}

type Option func(e *Error)

func New(options ...Option) *Error {
	err := &Error{}
	for _, v := range options {
		v(err)
	}
	return err
}

func WithCode(code int) Option {
	return func(e *Error) {
		e.Code = code
	}
}

func WithMsg(msg string) Option {
	return func(e *Error) {
		e.Msg = msg
	}
}

func WithData(data interface{}) Option {
	return func(e *Error) {
		e.Data = data
	}
}
