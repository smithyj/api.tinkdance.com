package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"tinkdance/application/app/api/internal/svc"
	"tinkdance/pkg/tracex"
)

var whiteList = map[string]struct{}{
	"/favicon.ico": {},
}

func Request(svcCtx *svc.ServiceContext) gin.HandlerFunc {
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

		// 请求上下文封装
		reqCtx := &svc.RequestContext{
			SvcCtx: svcCtx,
			Trace:  trace,
		}
		ctx := context.WithValue(context.Background(), svc.RequestKey, reqCtx)
		c.Set(svc.RequestKey, ctx)

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

		trace.AddAnnotation("host", c.Request.RemoteAddr)
		trace.AddAnnotation("method", c.Request.Method)
		trace.AddAnnotation("uri", c.Request.URL.Path)
		trace.AddAnnotation("query", c.Request.URL.RawQuery)
		trace.AddAnnotation("proto", c.Request.Proto)
		trace.AddAnnotation("status", fmt.Sprintf("%v", c.Writer.Status()))
		trace.AddAnnotation("user-agent", c.Request.UserAgent())
	}
}
