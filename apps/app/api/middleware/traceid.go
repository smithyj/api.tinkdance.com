package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"tinkdance/pkg/trace"
)

func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 TraceID 信息
		traceId := c.GetHeader(trace.TraceKey)
		if traceId == "" {
			cookie, err := c.Cookie(trace.TraceKey)
			if err == nil {
				traceId = cookie
			}
		}

		if traceId == "" {
			traceId = trace.NewTraceId()
		}

		// 输出 Header 头
		c.Header(trace.TraceKey, traceId)

		// 输出 Cookie 头
		host := strings.Split(c.Request.Host, ":")
		c.SetCookie(trace.TraceKey, traceId, 86400*365*100, "/", host[0], true, true)

		c.Set(trace.TraceKey, traceId)

		c.Next()
	}
}
