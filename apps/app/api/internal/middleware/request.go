package middleware

import (
	"context"

	"github.com/gin-gonic/gin"

	"tinkdance/apps/app/api/internal/svc"
	"tinkdance/pkg/trace"
)

func Request(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, ok := c.Get(trace.TraceKey)
		var t *trace.Trace
		if ok {
			t = v.(*trace.Trace)
		}
		// 请求上下文封装
		reqCtx := &svc.RequestContext{
			SvcCtx: svcCtx,
			Trace:  t,
		}
		ctx := context.WithValue(context.Background(), svc.RequestKey, reqCtx)
		c.Set(svc.RequestKey, ctx)

		c.Next()
	}
}
