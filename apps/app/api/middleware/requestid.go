package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"tinkdance/pkg/trace"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := trace.NewRequestId()
		// 输出 Header 头
		c.Header(trace.RequestKey, requestId)

		// 输出 Cookie 头
		host := strings.Split(c.Request.Host, ":")
		c.SetCookie(trace.RequestKey, requestId, 86400*365*100, "/", host[0], true, true)

		c.Set(trace.RequestKey, requestId)

		c.Next()
	}
}
