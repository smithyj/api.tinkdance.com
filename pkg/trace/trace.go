package trace

import (
	gouuid "github.com/satori/go.uuid"
)

const (
	TraceKey   = "X-Trace-Id"
	RequestKey = "X-Request-Id"
	initSpanId = "0"
)

func NewTraceId() string {
	return gouuid.NewV5(gouuid.NewV4(), "trace").String()
}

func NewRequestId() string {
	return gouuid.NewV5(gouuid.NewV4(), "request").String()
}
