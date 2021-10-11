package timeout

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CallBackFunc func(*http.Request)

type WhiteSkipFunc func(c *gin.Context) bool

type Option func(*Writer)

type Options struct {
	CallBack      CallBackFunc
	DefaultMsg    string
	Timeout       time.Duration
	ErrorHttpCode int
	WhiteSkipFunc WhiteSkipFunc
}

func WithTimeout(d time.Duration) Option {
	return func(t *Writer) {
		t.Timeout = d
	}
}

// WithErrorHttpCode Optional parameters
func WithErrorHttpCode(code int) Option {
	return func(t *Writer) {
		t.ErrorHttpCode = code
	}
}

// WithDefaultMsg Optional parameters
func WithDefaultMsg(s string) Option {
	return func(t *Writer) {
		t.DefaultMsg = s
	}
}

// WithCallBack Optional parameters
func WithCallBack(f CallBackFunc) Option {
	return func(t *Writer) {
		t.CallBack = f
	}
}

// WithWhiteSkip Optional parameters
func WithWhiteSkip(f WhiteSkipFunc) Option {
	return func(t *Writer) {
		t.WhiteSkipFunc = f
	}
}
