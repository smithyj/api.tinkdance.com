package middleware

import (
	"fmt"
	"net"
	"strings"

	"github.com/gin-gonic/gin"

	"tinkdance/application/app/api/internal/svc"
	"tinkdance/pkg/tracex"
)

func Logger(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	var whiteList = map[string]struct{}{
		"/favicon.ico": {},
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = fmt.Sprintf("%v?%v", path, raw)
		}

		// 获取 TraceID 信息
		key := c.GetHeader(tracex.TraceKey)
		if key == "" {
			cookie, err := c.Cookie(tracex.TraceKey)
			if err == nil {
				key = cookie
			}
		}

		trace := tracex.New(key)

		c.Set(tracex.TraceKey, trace)

		// 启动日志跟踪
		trace.Start()

		// 输出 Header 头
		c.Header(tracex.TraceKey, trace.TraceId())
		c.Header(tracex.RequestKey, trace.RequestId())

		// 输出 Cookie 头
		host := strings.Split(c.Request.Host, ":")
		c.SetCookie(tracex.TraceKey, trace.TraceId(), 86400*365*10, "/", host[0], true, false)

		c.Next()

		if _, ok := whiteList[c.Request.RequestURI]; ok {
			return
		}

		defer trace.Finish()

		ip := c.ClientIP()
		parseIP := net.ParseIP(ip)
		if parseIP.IsPrivate() || parseIP.IsLoopback() {
			ipHeaders := []string{"X-Real-IP", "X-Forwarded-For"}
			for _, v := range ipHeaders {
				if realIP := c.GetHeader(v); realIP != "" {
					ip = realIP
					break
				}
			}
		}

		trace.AddAnnotation("ip", ip)
		trace.AddAnnotation("method", c.Request.Method)
		trace.AddAnnotation("uri", c.Request.URL.Path)
		trace.AddAnnotation("query", c.Request.URL.RawQuery)
		trace.AddAnnotation("proto", c.Request.Proto)
		trace.AddAnnotation("status", fmt.Sprintf("%v", c.Writer.Status()))
		trace.AddAnnotation("user-agent", c.Request.UserAgent())
	}
}
