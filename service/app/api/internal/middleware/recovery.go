package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkdance/service/app/api/internal/svc"
	"tinkdance/service/pkg/errorx"
)

func Recovery(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *errorx.CodeError:
					c.JSON(http.StatusOK, err)
				default:
					e := err.(error)
					c.JSON(http.StatusInternalServerError, errorx.WithMsg(e.Error()))
				}
			}
		}()
		c.Next()
	}
}
