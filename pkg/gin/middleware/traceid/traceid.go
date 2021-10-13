package traceid

import (
	"github.com/gin-gonic/gin"
	gouuid "github.com/satori/go.uuid"
	"strings"
)

const TraceKey = "X-Trace-Id"

type trace struct {
	options options
}

func TraceId(opts ...Option) gin.HandlerFunc {
	req := trace{
		options: options{
			traceKey: TraceKey,
			idFunc: func() string {
				return NewTraceId()
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

		// 获取 TraceID 信息
		traceId := c.GetHeader(req.options.traceKey)
		if traceId == "" {
			cookie, err := c.Cookie(req.options.traceKey)
			if err == nil {
				traceId = cookie
			}
		}

		if traceId == "" {
			traceId = req.options.idFunc()
		}

		// 输出 Header 头
		c.Header(req.options.traceKey, traceId)

		// 输出 Cookie 头
		hostSlice := strings.Split(c.Request.Host, ":")

		host := hostSlice[0]
		if req.options.cookie.Host != "" {
			host = req.options.cookie.Host
		}

		c.SetCookie(req.options.traceKey, traceId, req.options.cookie.MaxAge, req.options.cookie.Path, host, req.options.cookie.Secure, req.options.cookie.HttpOnly)

		c.Set(req.options.traceKey, traceId)

		c.Next()
	}
}

func NewTraceId() string {
	return gouuid.NewV5(gouuid.NewV4(), "trace").String()
}
