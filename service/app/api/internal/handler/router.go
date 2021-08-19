package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tinkdance/pkg/tinkgo/httpx"
	"tinkdance/service/app/api/internal/handler/account"
	"tinkdance/service/app/api/internal/handler/passport"
	"tinkdance/service/app/api/internal/middleware"
	"tinkdance/service/app/api/internal/svc"
)

func NewRouter(server *httpx.Server, svcCtx *svc.ServiceContext) {
	engine := server.Engine()
	// 全局中间件
	{
		engine.Use(middleware.Logger(svcCtx))
		engine.Use(gin.Recovery())
		engine.Use(cors.Default())
		engine.Use(middleware.Recovery(svcCtx))
	}
	// 通行证
	{
		g := engine.Group("/passport")
		g.POST("/register", passport.RegisterHandler(svcCtx))
		g.POST("/login", passport.LoginHandler(svcCtx))
		g.POST("/logout", passport.LogoutHandler(svcCtx))
	}
	// 账号
	{
		g := engine.Group("/account")
		g.GET("/:id", account.ProfileHandler(svcCtx))
		g.PUT("/:id", account.UpdateProfileHandler(svcCtx))
	}
}
