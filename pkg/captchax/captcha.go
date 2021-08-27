package captchax

import (
	"context"
	"strings"
	"time"
	"tinkdance/pkg/redisx"
)

const key = "captcha:"
const ttl = time.Minute * 10

type Captcha interface {
	Set(ctx context.Context, captchaId, code string) bool
	Validate(ctx context.Context, captchaId, code string) bool
}

type captcha struct {
	redis redisx.Redis
}

func (c *captcha) Set(ctx context.Context, captchaId, code string) bool {
	cmd := c.redis.Set(ctx, key+captchaId, code, ttl)
	if err := cmd.Err(); err != nil {
		return false
	}
	return true
}

func (c *captcha) Validate(ctx context.Context, captchaId, code string) bool {
	cmd := c.redis.Get(ctx, key+captchaId)
	if err := cmd.Err(); err != nil {
		c.Del(ctx, captchaId)
		return false
	}
	if strings.ToLower(cmd.Val()) != strings.ToLower(code) {
		return false
	}
	return true
}

func (c *captcha) Del(ctx context.Context, captchaId string) bool {
	cmd := c.redis.Del(ctx, key+captchaId)
	if err := cmd.Err(); err != nil {
		return false
	}
	return true
}

func New(redis redisx.Redis) Captcha {
	return &captcha{
		redis: redis,
	}
}
