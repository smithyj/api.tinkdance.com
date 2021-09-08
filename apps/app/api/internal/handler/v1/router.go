package v1

import (
	"github.com/gin-gonic/gin"
	"tinkdance/apps/app/api/internal/handler/v1/captcha"
)

func Router(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		// 验证码
		g := v1.Group("/captcha")
		// 图片验证码
		g.GET("/image", captcha.Image())
	}
}
