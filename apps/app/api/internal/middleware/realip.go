package middleware

import (
	"github.com/gin-gonic/gin"
)

func RealIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.RemoteAddr = "223.5.5.5:39392"
		c.Next()
	}
}
