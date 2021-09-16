package captcha

import (
	"context"
	"strconv"
	"testing"
	"time"
	"tinkdance/pkg/redis"
)

func TestCaptcha(t *testing.T)  {
	rdb, err := redis.New(redis.WithConfig(redis.Config{
		Addr: "127.0.0.1",
		Port: 6379,
		Password: "",
		DB: 0,
		Prefix: "tinkdance",
	}))
	if err != nil {
		t.FailNow()
	}
	captcha := New(WithRedis(rdb))
	ctx := context.Background()
	key := "test"
	value := time.Now().Unix()
	ok := captcha.Storage(ctx, key, strconv.Itoa(int(value)))
	if !ok {
		t.FailNow()
	}
	if ok := captcha.Validate(ctx, key, strconv.Itoa(int(value))); !ok {
		t.FailNow()
	}
}
