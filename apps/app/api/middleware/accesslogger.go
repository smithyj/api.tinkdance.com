package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tinkdance/apps/app/api/pkg/server"
)

func AccessLogger(srvCtx *server.Context) gin.HandlerFunc {
	blackList := map[string]struct{} {
		"/favicon.ico": {},
	}

	return func(c *gin.Context) {
		c.Next()

		path := c.Request.URL.Path

		if _, ok := blackList[path]; ok {
			return
		}

		srvCtx.AccessLogger.WithOptions(zap.WithCaller(false)).With(
			zap.String("remote", c.Request.RemoteAddr),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("uri", c.Request.RequestURI),
			zap.String("host", c.Request.Host),
			zap.Int("status", c.Writer.Status()),
			zap.Int("length", c.Writer.Size()),
			zap.String("proto", c.Request.Proto),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("referer", c.Request.Referer()),
		).Info("")
	}
}
