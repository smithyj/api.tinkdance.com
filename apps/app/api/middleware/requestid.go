package middleware

import (
	"github.com/gin-gonic/gin"
	"tinkdance/pkg/gin/middleware/requestid"
	"tinkdance/pkg/trace"
)

func RequestId() gin.HandlerFunc {
	return requestid.RequestId(
		requestid.WithRequestKey(trace.RequestKey),
		requestid.WithIdFunc(trace.NewRequestId),
	)
}
