package captcha

import (
	"context"
	"fmt"
	"strings"
	"time"

	"thinkdance/pkg/redis"
)

// TTL 验证码超时时间
const TTL = 10 * time.Minute

type Captcha interface {
	Storage(ctx context.Context, captchaId, code string) bool
	Validate(ctx context.Context, captchaId, code string) bool
	i()
}

type captcha struct {
	redis redis.Redis
}

func (c *captcha) prefix(captchaId string) string {
	return fmt.Sprintf("captcha:%s", captchaId)
}

func (c *captcha) Storage(ctx context.Context, captchaId, code string) bool {
	key := c.prefix(captchaId)
	expiration := TTL
	if err := c.redis.Set(ctx, key, code, expiration).Err(); err != nil {
		return false
	}
	return true
}

func (c *captcha) Validate(ctx context.Context, captchaId, code string) bool {
	key := c.prefix(captchaId)
	v, err := c.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		return false
	}

	if strings.ToLower(v) != strings.ToLower(code) {
		return false
	}

	// 验证成功，删除验证码
	c.redis.Del(ctx, key)

	return true
}

func (c *captcha) i() {}

func NewCaptcha(redis redis.Redis) Captcha {
	return &captcha{
		redis: redis,
	}
}
