package cache

import (
	"context"
	"fmt"
	"time"

	"tinkdance/pkg/redis"
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, bool)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
	Del(ctx context.Context, keys ...string) bool
	i()
}

type cache struct {
	redis redis.Redis
}

func (c *cache) prefix(key string) string {
	return fmt.Sprintf("cache:%s", key)
}

func (c *cache) Get(ctx context.Context, key string) ([]byte, bool) {
	key = c.prefix(key)
	v, err := c.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, false
	} else if err != nil {
		return nil, false
	}
	return []byte(v), true
}

func (c *cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	key = c.prefix(key)
	if err := c.redis.Set(ctx, key, value, expiration).Err(); err != nil {
		return false
	}
	return true
}

func (c *cache) Del(ctx context.Context, keys ...string) bool {
	for index, value := range keys {
		keys[index] = c.prefix(value)
	}
	if err := c.redis.Del(ctx, keys...).Err(); err != nil {
		return false
	}
	return true
}

func (c *cache) i() {}

func NewCache(redis redis.Redis) Cache {
	return &cache{
		redis: redis,
	}
}
