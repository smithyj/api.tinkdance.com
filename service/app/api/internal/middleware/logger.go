package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"tinkdance/pkg/tinkgo/tracex"
	"tinkdance/service/app/api/internal/svc"
)

func ignoreCheck(str string) bool {
	whites := []string{
		"/favicon.ico",
	}
	for _, v := range whites {
		if v == str {
			return true
		}
	}
	return false
}

func Logger(srvCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		startAt := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = fmt.Sprintf("%v?%v", path, raw)
		}

		t := tracex.New(c.GetHeader(tracex.TracingKey))
		t.Start()

		c.Set(tracex.TracingKey, t)

		c.Next()

		endAt := time.Now()

		if !ignoreCheck(c.Request.RequestURI) {
			zap.L().WithOptions(zap.WithCaller(false)).Info(
				"[REQUEST]",
				zap.String(tracex.TracingKey, t.TraceId()),
				zap.String("request_id", t.RequestId()),
				zap.String("host", c.Request.RemoteAddr),
				zap.String("method", c.Request.Method),
				zap.String("uri", c.Request.RequestURI),
				zap.String("proto", c.Request.Proto),
				zap.String("status", fmt.Sprintf("%v", c.Writer.Status())),
				zap.Duration("duration", endAt.Sub(startAt)),
				zap.String("user-agent", c.Request.UserAgent()),
			)
		}
	}
}
