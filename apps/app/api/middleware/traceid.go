package middleware

import (
	"github.com/gin-gonic/gin"
	"tinkdance/pkg/gin/middleware/traceid"
	"tinkdance/pkg/trace"
)

func TraceId() gin.HandlerFunc {
	return traceid.TraceId(
		traceid.WithTraceKey(trace.TraceKey),
		traceid.WithIdFunc(trace.NewTraceId),
	)
}
