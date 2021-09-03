package handler

import (
	"tinkdance/application/app/api/internal/handler/v1"
	middleware2 "tinkdance/application/app/api/internal/middleware"
	"tinkdance/application/app/api/internal/svc"
	"tinkdance/pkg/httpx"
)

func Router(server *httpx.Server, svcCtx *svc.ServiceContext) {
	engine := server.Engine()
	// 全局中间件
	{
		engine.Use(middleware2.SameSite())
		engine.Use(middleware2.Cors(svcCtx))
		engine.Use(middleware2.Logger(svcCtx))
		engine.Use(middleware2.Request(svcCtx))
		engine.Use(middleware2.Recovery())
	}
	v1.Router(engine)
}
