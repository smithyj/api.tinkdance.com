package middleware

import (
	"fmt"
	"net"
	"strings"

	"github.com/gin-gonic/gin"

	"tinkdance/apps/app/api/internal/svc"
	"tinkdance/pkg/trace"
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
		key := c.GetHeader(trace.TraceKey)
		if key == "" {
			cookie, err := c.Cookie(trace.TraceKey)
			if err == nil {
				key = cookie
			}
		}

		t := trace.New(key)

		c.Set(trace.TraceKey, t)

		// 启动日志跟踪
		t.Start()

		// 输出 Header 头
		c.Header(trace.TraceKey, t.TraceId())
		c.Header(trace.RequestKey, t.RequestId())

		// 输出 Cookie 头
		host := strings.Split(c.Request.Host, ":")
		c.SetCookie(trace.TraceKey, t.TraceId(), 86400*365*10, "/", host[0], true, false)

		c.Next()

		if _, ok := whiteList[c.Request.RequestURI]; ok {
			return
		}

		defer t.Finish()

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

		t.AddAnnotation("ip", ip)
		t.AddAnnotation("method", c.Request.Method)
		t.AddAnnotation("uri", c.Request.URL.Path)
		t.AddAnnotation("query", c.Request.URL.RawQuery)
		t.AddAnnotation("proto", c.Request.Proto)
		t.AddAnnotation("status", fmt.Sprintf("%v", c.Writer.Status()))
		t.AddAnnotation("user-agent", c.Request.UserAgent())
	}
}
