package captcha

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkdance/apps/app/api/internal/logic/captcha"
	"tinkdance/apps/app/api/internal/svc"
)

func Image() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req captcha.ImageRequest

		ctx, _, err := svc.RequestWithGin(c)
		if err != nil {
			panic(err)
		}

		resp, err := captcha.Image(ctx, &req)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, resp)
	}
}
