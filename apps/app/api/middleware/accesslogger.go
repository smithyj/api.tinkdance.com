package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"time"
	"tinkdance/apps/app/api/pkg/server"
	"tinkdance/pkg/trace"
)

func AccessLogger(srvCtx *server.Context) gin.HandlerFunc {
	blackList := map[string]struct{} {
		"/favicon.ico": {},
	}

	return func(c *gin.Context) {

		startAt := time.Now()

		c.Next()

		path := c.Request.URL.Path

		if _, ok := blackList[path]; ok {
			return
		}

		traceId := ""
		if v, ok := c.Get(trace.TraceKey); ok {
			traceId = v.(string)
		}

		requestId := ""
		if v, ok := c.Get(trace.RequestKey); ok {
			requestId = v.(string)
		}

		host, port, _ := net.SplitHostPort(c.Request.RemoteAddr)

		srvCtx.AccessLogger.WithOptions(zap.WithCaller(false)).With(
			zap.String("trace_id", traceId),
			zap.String("request_id", requestId),
			zap.String("ip", host),
			zap.String("port", port),
			zap.String("remote", c.Request.RemoteAddr),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("host", c.Request.Host),
			zap.Int("status", c.Writer.Status()),
			zap.Int("length", c.Writer.Size()),
			zap.String("proto", c.Request.Proto),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("referer", c.Request.Referer()),
			zap.Int64("elapsed", time.Now().Sub(startAt).Microseconds()),
		).Info("")
	}
}
