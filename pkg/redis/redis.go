package redis

import (
	"context"
	"fmt"
	goredis "github.com/go-redis/redis/v8"
	"time"
)

type Redis interface {
	Get(ctx context.Context, key string) *goredis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.StatusCmd
	Del(ctx context.Context, keys ...string) *goredis.IntCmd
	i()
}

type redisx struct {
	client *goredis.Client
	prefix string
}

func (r *redisx) build(key string) string {
	if r.prefix == "" {
		return key
	}
	return fmt.Sprintf("%v:%v", r.prefix, key)
}

func (r *redisx) Get(ctx context.Context, key string) *goredis.StringCmd {
	return r.client.Get(ctx, r.build(key))
}

func (r *redisx) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.StatusCmd {
	return r.client.Set(ctx, r.build(key), value, expiration)
}

func (r *redisx) Del(ctx context.Context, keys ...string) *goredis.IntCmd {
	for index, key := range keys {
		keys[index] = r.build(key)
	}
	return r.client.Del(ctx, keys...)
}

func (r *redisx) i() {}

func New(cfg Config) (Redis, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}
	return &redisx{
		client: client,
		prefix: cfg.Prefix,
	}, nil
}
