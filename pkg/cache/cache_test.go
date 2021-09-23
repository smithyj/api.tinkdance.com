package cache

import (
	"context"
	"testing"
	"time"
	"tinkdance/pkg/redis"
)

func TestCache(t *testing.T) {
	rdb, err := redis.New(redis.WithConfig(redis.Config{
		Addr:     "127.0.0.1",
		Port:     6379,
		Password: "",
		DB:       0,
		Prefix:   "tinkdance",
	}))
	if err != nil {
		t.FailNow()
	}
	cache := New(WithRedis(rdb))
	ctx := context.Background()
	key := "test"
	value := time.Now().Unix()
	ok := cache.Set(ctx, key, value, redis.KeepTTL)
	if !ok {
		t.FailNow()
	}
	_, ok = cache.Get(ctx, key)
	if !ok {
		t.FailNow()
	}
}
