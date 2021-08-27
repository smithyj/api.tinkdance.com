package tracex

import (
	"context"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
	"time"
	"tinkdance/pkg/logger"
)

const (
	TraceKey   = "X-Trace-Id"
	RequestKey = "X-Request-Id"
	initSpanId = "0"
)

type Trace struct {
	logger       *zerolog.Event
	traceId      string
	requestId    string
	spanId       string
	parentSpanId string
	startAt      time.Time
	endAt        time.Time
	annotation   map[string]string
}

func (t *Trace) AddAnnotation(key string, value string) *Trace {
	t.annotation[key] = value
	return t
}

func (t *Trace) TraceId() string {
	return t.traceId
}

func (t *Trace) RequestId() string {
	return t.requestId
}

func (t *Trace) SpanId() string {
	return t.spanId
}

func (t *Trace) ParentSpanId() string {
	return t.parentSpanId
}

func (t *Trace) Annotation() map[string]string {
	return t.annotation
}

func (t *Trace) Start() *Trace {
	t.startAt = time.Now()
	return t
}

func (t *Trace) Finish() {
	if t.logger == nil {
		return
	}
	t.endAt = time.Now()

	t.logger.Fields(map[string]interface{}{
		"trace_id":       t.traceId,
		"request_id":     t.requestId,
		"span_id":        t.spanId,
		"parent_span_id": t.parentSpanId,
		"annotation":     t.annotation,
		"duration":       t.endAt.Sub(t.startAt),
	}).Msg("")
}

func (t *Trace) Fork() *Trace {
	spanId := uuid.NewV5(uuid.NewV4(), "span-log").String()
	return &Trace{
		logger:       t.logger,
		traceId:      t.traceId,
		spanId:       spanId,
		parentSpanId: t.spanId,
		requestId:    t.requestId,
		startAt:      time.Now(),
		annotation:   make(map[string]string),
	}
}

func New(traceId string) *Trace {
	if traceId == "" {
		traceId = uuid.NewV5(uuid.NewV4(), "trace-log").String()
	}
	requestId := uuid.NewV5(uuid.NewV4(), "request-log").String()
	return &Trace{
		logger:       logger.Trace(),
		traceId:      traceId,
		requestId:    requestId,
		spanId:       initSpanId,
		parentSpanId: initSpanId,
		annotation:   make(map[string]string),
	}
}

func WithContext(ctx context.Context) *Trace {
	v := ctx.Value(TraceKey)
	t, ok := v.(*Trace)
	if !ok {
		t = New("")
	}
	return t
}
