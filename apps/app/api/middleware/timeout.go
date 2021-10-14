package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"tinkdance/pkg/code/bizcode"
	"tinkdance/pkg/error/bizerror"
	"tinkdance/pkg/gin/middleware/timeout"
)


func Timeout(duration time.Duration) gin.HandlerFunc {
	var whiteList = map[string]struct{}{}

	code := bizcode.RequestTimeout
	return timeout.Timeout(
		timeout.WithWhiteSkip(func(c *gin.Context) bool {
			if _, ok := whiteList[c.Request.RequestURI]; ok {
				return true
			}
			return false
		}),
		timeout.WithTimeout(duration),
		timeout.WithDefaultMsg(`{"code":` + strconv.Itoa(code) + `,"msg":"`+ bizerror.New(bizerror.WithCode(code)).Format().Msg +`"}`),
	)
}
