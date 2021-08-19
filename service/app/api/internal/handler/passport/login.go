package passport

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkdance/service/app/api/internal/logic/passport"
	"tinkdance/service/app/api/internal/svc"
)

func LoginHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req passport.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}
		l := passport.NewLoginLogic(context.TODO(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, resp)
	}
}