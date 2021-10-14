package captcha

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkdance/apps/app/api/logic/captcha"
)

func Image() gin.HandlerFunc {
	return func(c *gin.Context) {
		//time.Sleep(10 * time.Second)
		resp := captcha.Image(context.TODO(), &captcha.ImageRequest{})
		c.JSON(http.StatusOK, resp)
	}
}
