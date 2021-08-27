package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkdance/pkg/errorx"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *errorx.Error:
					e := err.(*errorx.Error)
					c.JSON(http.StatusOK, e)
				case string:
					e := err.(string)
					c.JSON(http.StatusInternalServerError, errorx.WithMsg(e))
				default:
					e := err.(error)
					c.JSON(http.StatusInternalServerError, errorx.WithMsg(e.Error()))
				}
			}
		}()
		c.Next()
	}
}
