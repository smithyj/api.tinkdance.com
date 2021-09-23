package server

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"tinkdance/apps/app/api/config"
	"tinkdance/pkg/cache"
	"tinkdance/pkg/captcha"
	"tinkdance/pkg/database"
	"tinkdance/pkg/env"
	"tinkdance/pkg/logger"
	"tinkdance/pkg/redis"
)

type Option func(ctx *Context)

type Context struct {
	// 系统配置
	Config *config.Config
	// 访问日志
	AccessLogger *zap.Logger
	// 链路日志
	TraceLogger *zap.Logger
	// 数据库组件
	Database *gorm.DB
	// Redis 组件
	Redis redis.Redis
	// 缓存组件
	Cache cache.Cache
	// 验证码组件
	Captcha captcha.Captcha
}

func WithConfig(cfg *config.Config) Option {
	return func(ctx *Context) {
		ctx.Config = cfg
	}
}

func NewServerContext(options ...Option) *Context {
	var ctx = new(Context)

	for _, f := range options {
		f(ctx)
	}

	if ctx.Config == nil {
		panic(errors.New("缺少配置信息"))
	}

	// 初始化 Access Logger
	ctx.AccessLogger = logger.New(
		logger.WithField("project", fmt.Sprintf("%s[%s]", config.Name, env.Active().Value())),
		logger.WithDebugLevel(ctx.Config.Debug),
		logger.WithDisableConsole(!env.Active().IsDev()),
		logger.WithFileRotationP(config.AccessLogFile),
	)

	// 初始化 Trace Logger
	ctx.TraceLogger = logger.New(
		logger.WithField("project", fmt.Sprintf("%s[%s]", config.Name, env.Active().Value())),
		logger.WithDebugLevel(ctx.Config.Debug),
		logger.WithDisableConsole(!env.Active().IsDev()),
		logger.WithFileRotationP(config.TraceLogFile),
	)

	// 初始化数据库
	db, err := database.New(database.WithConfig(ctx.Config.Database))
	if err != nil {
		panic(err)
	}
	ctx.Database = db

	// 初始化 Redis
	rdb, err := redis.New(redis.WithConfig(ctx.Config.Redis))
	if err != nil {
		panic(err)
	}
	ctx.Redis = rdb

	// 初始化缓存
	ctx.Cache = cache.New(cache.WithRedis(rdb))

	// 初始化验证码
	ctx.Captcha = captcha.New(captcha.WithRedis(rdb))

	return ctx
}
