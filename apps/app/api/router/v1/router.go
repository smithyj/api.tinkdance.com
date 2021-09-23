package v1

import (
	"github.com/gin-gonic/gin"
	"tinkdance/apps/app/api/handler/captcha"
	"tinkdance/apps/app/api/pkg/server"
)

func Run(srvCtx *server.Context, engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("/captcha/image", captcha.Image())
	}
}
