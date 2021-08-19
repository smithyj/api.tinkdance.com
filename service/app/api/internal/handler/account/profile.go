package account

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkdance/service/app/api/internal/logic/account"
	"tinkdance/service/app/api/internal/svc"
)

func ProfileHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req account.ProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}
		l := account.NewProfileLogic(context.TODO(), svcCtx)
		resp, err := l.Profile(&req)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, resp)
	}
}
