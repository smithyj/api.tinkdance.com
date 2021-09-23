package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkdance/pkg/error/bizerror"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var resp *bizerror.Error
				switch err.(type) {
				case string, int, int8, int16, int32, int64, float32, float64:
					resp = bizerror.New(bizerror.WithMsg(fmt.Sprintf("%v", err)))
				case *string:
					resp = bizerror.New(bizerror.WithMsg(fmt.Sprintf("%v", *err.(*string))))
				case *int:
					resp = bizerror.New(bizerror.WithMsg(fmt.Sprintf("%v", *err.(*int))))
				case *int8:
					resp = bizerror.New(bizerror.WithMsg(fmt.Sprintf("%v", *err.(*int8))))
				case *int16:
					resp = bizerror.New(bizerror.WithMsg(fmt.Sprintf("%v", *err.(*int16))))
				case *int32:
					resp = bizerror.New(bizerror.WithMsg(fmt.Sprintf("%v", *err.(*int32))))
				case *int64:
					resp = bizerror.New(bizerror.WithMsg(fmt.Sprintf("%v", *err.(*int64))))
				case bizerror.Error:
					e := err.(bizerror.Error)
					resp = &e
				case *bizerror.Error:
					e := err.(*bizerror.Error)
					resp = e
				default:
					e := err.(error)
					resp = bizerror.New(bizerror.WithMsg(e.Error()))
				}
				c.JSON(http.StatusInternalServerError, resp.Format())
			}
		}()
		c.Next()
	}
}
