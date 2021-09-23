package main

import (
	"github.com/gin-gonic/gin"
	"tinkdance/apps/app/api/config"
	"tinkdance/apps/app/api/pkg/server"
	"tinkdance/apps/app/api/router"
	"tinkdance/pkg/env"
)

func main() {
	// 初始化资源
	srvCtx := server.NewServerContext(
		server.WithConfig(config.Setup()),
	)

	// 正式环境，关闭 Gin Debug 模式
	if env.Active().IsProd() {
		gin.SetMode("release")
	}

	engine := gin.New()

	router.Run(srvCtx, engine)

	_ = engine.Run(":8080")
}
