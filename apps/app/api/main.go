package main

import (
	"flag"

	"tinkdance/apps/app/api/internal/config"
	"tinkdance/apps/app/api/internal/handler"
	"tinkdance/apps/app/api/internal/svc"
	"tinkdance/pkg/httpx"
	"tinkdance/pkg/logger"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "dev", "设置环境变量，默认 dev")
	flag.Parse()

	// 全局配置
	conf, err := config.NewConfig(env)
	if err != nil {
		panic(err)
	}

	// 日志初始化
	logger.Setup(conf.Log)

	// 上下文初始化
	svcCtx, err := svc.NewServiceContext(&conf)
	if err != nil {
		logger.Error().Msg(err.Error())
		return
	}

	// 服务初始化
	srv := httpx.NewServer(conf.Mode)

	// 路由初始化
	handler.Router(srv, svcCtx)

	// 服务运行
	srv.GraceRun(conf.Addr)
}
