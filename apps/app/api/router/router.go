package router

import (
	"github.com/gin-gonic/gin"
	"time"
	"tinkdance/apps/app/api/middleware"
	"tinkdance/apps/app/api/pkg/server"
	"tinkdance/apps/app/api/router/v1"
)

func Run(srvCtx *server.Context, engine *gin.Engine) {
	{
		// 全局中间件
		engine.Use(middleware.TraceId())
		engine.Use(middleware.RequestId())
		engine.Use(middleware.AccessLogger(srvCtx))
		engine.Use(middleware.Timeout(3 * time.Second))
		engine.Use(middleware.Recovery())
	}
	// v1 路由
	v1.Run(srvCtx, engine)
}
