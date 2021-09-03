package middleware

import (
	"context"

	"github.com/gin-gonic/gin"

	"tinkdance/application/app/api/internal/svc"
	"tinkdance/pkg/tracex"
)

func Request(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, ok := c.Get(tracex.TraceKey)
		var trace *tracex.Trace
		if ok {
			trace = v.(*tracex.Trace)
		}
		// 请求上下文封装
		reqCtx := &svc.RequestContext{
			SvcCtx: svcCtx,
			Trace:  trace,
		}
		ctx := context.WithValue(context.Background(), svc.RequestKey, reqCtx)
		c.Set(svc.RequestKey, ctx)

		c.Next()
	}
}
