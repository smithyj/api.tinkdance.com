package svc

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"tinkdance/pkg/tracex"
)

const RequestKey = "request-context"

type RequestContext struct {
	SvcCtx *ServiceContext
	Trace  *tracex.Trace
}

// RequestWithContext 从上下文中获取 RequestContext
func RequestWithContext(ctx context.Context) (context.Context, *RequestContext, error) {
	v, ok := ctx.Value(RequestKey).(*RequestContext)
	if !ok {
		return ctx, nil, errors.New("RequestContext is not initialize")
	}
	return ctx, v, nil
}

func RequestWithGin(c *gin.Context) (context.Context, *RequestContext, error) {
	v, ok := c.Get(RequestKey)
	if !ok {
		return context.Background(), nil, errors.New("RequestContext is not initialize")
	}

	ctx, ok := v.(context.Context)
	if !ok {
		return context.Background(), nil, errors.New("RequestContext is initialize error")
	}

	return RequestWithContext(ctx)
}
