package handler

import (
	"tinkdance/application/app/api/internal/handler/middleware"
	"tinkdance/application/app/api/internal/handler/v1"
	"tinkdance/application/app/api/internal/svc"
	"tinkdance/pkg/httpx"
)

func Router(server *httpx.Server, svcCtx *svc.ServiceContext) {
	engine := server.Engine()
	// 全局中间件
	{
		engine.Use(middleware.SameSite())
		engine.Use(middleware.Cors(svcCtx))
		engine.Use(middleware.Request(svcCtx))
		engine.Use(middleware.Recovery())
	}
	v1.Router(engine)
}
