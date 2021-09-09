package captcha

import (
	"context"
	"strconv"
	"testing"
	"thinkdance/pkg/redis"
	"time"
)

func TestCaptcha(t *testing.T)  {
	rdb, err := redis.NewRedis(redis.Config{
		Addr: "127.0.0.1",
		Port: 6379,
		Password: "",
		DB: 0,
		Prefix: "tinkdance",
	})
	if err != nil {
		t.FailNow()
	}
	captcha := NewCaptcha(rdb)
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
