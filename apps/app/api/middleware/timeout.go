package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vearne/gin-timeout"
	"strconv"
	"time"
	"tinkdance/pkg/code/bizcode"
	"tinkdance/pkg/error/bizerror"
)

func Timeout(duration time.Duration) gin.HandlerFunc {
	code := bizcode.RequestTimeout
	return timeout.Timeout(
		timeout.WithTimeout(duration),
		timeout.WithDefaultMsg(`{"code":` + strconv.Itoa(code) + `,"msg":"`+ bizerror.New(bizerror.WithCode(code)).Format().Msg +`"}`),
	)
}
