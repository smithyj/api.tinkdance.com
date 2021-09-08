package handler

import (
	"tinkdance/apps/app/api/internal/handler/v1"
	"tinkdance/apps/app/api/internal/middleware"
	"tinkdance/apps/app/api/internal/svc"
	"tinkdance/pkg/httpx"
)

func Router(server *httpx.Server, svcCtx *svc.ServiceContext) {
	engine := server.Engine()
	// 全局中间件
	{
		engine.Use(middleware.RealIP())
		engine.Use(middleware.SameSite())
		engine.Use(middleware.Cors(svcCtx))
		engine.Use(middleware.Logger(svcCtx))
		engine.Use(middleware.Request(svcCtx))
		engine.Use(middleware.Recovery())
	}
	v1.Router(engine)
}
