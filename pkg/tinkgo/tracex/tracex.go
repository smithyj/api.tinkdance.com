package tracex

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

const (
	TracingKey = "trace_id"
	initSpanId = "0"
)

type Trace struct {
	logger       *zap.Logger
	traceId      string
	spanId       string
	parentSpanId string
	requestId    string
	annotation   map[string]interface{}
	startAt      time.Time
	endAt        time.Time
}

func (t *Trace) setLogger(logger *zap.Logger) *Trace {
	t.logger = logger
	return t
}

func (t *Trace) setTraceId(traceId string) *Trace {
	t.traceId = traceId
	return t
}

func (t *Trace) setSpanId(spanId string) *Trace {
	t.spanId = spanId
	return t
}

func (t *Trace) setParentSpanId(parentSpanId string) *Trace {
	t.parentSpanId = parentSpanId
	return t
}

func (t *Trace) SetAnnotation(key string, value interface{}) *Trace {
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

func (t *Trace) Annotation() map[string]interface{} {
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
	var fields = make([]zap.Field, 0)
	fields = append(fields, zap.String("trace_id", t.traceId))
	fields = append(fields, zap.String("request_id", t.requestId))
	fields = append(fields, zap.String("span_id", t.spanId))
	fields = append(fields, zap.String("parent_span_id", t.parentSpanId))
	if len(t.annotation) > 0 {
		fields = append(fields, zap.Any("annotation", t.annotation))
	}
	fields = append(fields, zap.Duration("duration", t.endAt.Sub(t.startAt)))

	t.logger.WithOptions(zap.AddCallerSkip(1)).Info(
		"trace",
		fields...,
	)
}

func (t *Trace) Fork() *Trace {
	spanId := uuid.NewV5(uuid.NewV4(), "span").String()
	return &Trace{
		logger:       t.logger,
		traceId:      t.traceId,
		spanId:       spanId,
		parentSpanId: t.spanId,
		requestId:    t.requestId,
		startAt:      time.Now(),
		annotation:   make(map[string]interface{}),
	}
}

func New(traceId string) *Trace {
	if traceId == "" {
		traceId = uuid.NewV5(uuid.NewV4(), "trace").String()
	}
	requestId := uuid.NewV5(uuid.NewV4(), "request").String()
	return &Trace{
		logger:       zap.L(),
		traceId:      traceId,
		spanId:       initSpanId,
		parentSpanId: initSpanId,
		requestId:    requestId,
		annotation:   make(map[string]interface{}),
	}
}

func WithContext(ctx context.Context) *Trace {
	v := ctx.Value(TracingKey)
	t, ok := v.(*Trace)
	if !ok {
		t = New("")
	}
	return t
}
