package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"tinkdance/pkg/tinkgo/httpx"
	"tinkdance/pkg/tinkgo/logx"
	"tinkdance/service/app/api/internal/config"
	"tinkdance/service/app/api/internal/handler"
	"tinkdance/service/app/api/internal/svc"
)

var env = kingpin.Flag("env", "Set run environment, options: prod / test / dev").Default("dev").String()

func main() {
	// 命令行解析
	kingpin.Version("1.0.0")
	kingpin.Parse()

	// 全局配置
	conf, err := config.NewConfig(*env)
	if err != nil {
		panic(err)
	}

	// 日志初始化
	logx.Setup(conf.LogConfig)

	// 上下文初始化
	srvCtx, err := svc.NewServiceContext(conf)
	if err != nil {
		panic(err)
	}

	// 服务初始化
	server := httpx.NewServer(conf.Mode)

	// 路由初始化
	handler.NewRouter(server, srvCtx)

	// 服务运行
	server.GraceRun(conf.Addr)
}
