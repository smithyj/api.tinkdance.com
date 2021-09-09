package redis

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

const (
	KeepTTL = goredis.KeepTTL
	Nil = goredis.Nil
)

type Redis interface {
	Get(ctx context.Context, key string) *goredis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.StatusCmd
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.BoolCmd
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.BoolCmd
	Del(ctx context.Context, keys ...string) *goredis.IntCmd
	i()
}

type redis struct {
	client *goredis.Client
	config Config
}

func (r *redis) prefix(key string) string {
	if r.config.Prefix != "" {
		return fmt.Sprintf("%s:%s", r.config.Prefix, key)
	}
	return key
}

func (r *redis) Get(ctx context.Context, key string) *goredis.StringCmd {
	key = r.prefix(key)
	return r.client.Get(ctx, key)
}

func (r *redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.StatusCmd {
	key = r.prefix(key)
	return r.client.Set(ctx, key, value, expiration)
}

func (r *redis) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.StatusCmd {
	key = r.prefix(key)
	return r.client.SetEX(ctx, key, value, expiration)
}

func (r *redis) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.BoolCmd {
	key = r.prefix(key)
	return r.client.SetNX(ctx, key, value, expiration)
}

func (r *redis) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *goredis.BoolCmd {
	key = r.prefix(key)
	return r.client.SetXX(ctx, key, value, expiration)
}

func (r *redis) Del(ctx context.Context, keys ...string) *goredis.IntCmd {
	for index, value := range keys {
		keys[index] = r.prefix(value)
	}
	return r.client.Del(ctx, keys...)
}

func (r *redis) i() {}

func NewRedis(config Config) (Redis, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Addr, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}
	return &redis{
		client: client,
		config: config,
	}, nil
}
