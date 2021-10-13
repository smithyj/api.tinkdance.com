package requestid

import (
	"github.com/gin-gonic/gin"
	gouuid "github.com/satori/go.uuid"
	"strings"
)

const RequestKey = "X-Request-Id"

type request struct {
	options options
}

func RequestId(opts ...Option) gin.HandlerFunc {
	req := request{
		options: options{
			requestKey: RequestKey,
			idFunc: func() string {
				return NewRequestId()
			},
			cookie: Cookie{
				MaxAge: 86400 * 365 * 100,
				Path: "/",
				Host: "",
				Secure: true,
				HttpOnly: true,
			},
		},
	}
	return func(c *gin.Context) {
		for _, opt := range opts {
			opt(&req)
		}
		requestId := req.options.idFunc()
		// 输出 Header 头
		c.Header(req.options.requestKey, requestId)

		// 输出 Cookie 头
		hostSlice := strings.Split(c.Request.Host, ":")

		host := hostSlice[0]
		if req.options.cookie.Host != "" {
			host = req.options.cookie.Host
		}

		c.SetCookie(req.options.requestKey, requestId, req.options.cookie.MaxAge, req.options.cookie.Path, host, req.options.cookie.Secure, req.options.cookie.HttpOnly)

		c.Set(req.options.requestKey, requestId)

		c.Next()
	}
}

func NewRequestId() string {
	return gouuid.NewV5(gouuid.NewV4(), "request").String()
}
