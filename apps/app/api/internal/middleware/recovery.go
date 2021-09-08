package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkdance/pkg/errorx"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case string:
					e := err.(string)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(e))
				case *string:
					e := err.(*string)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(*e))
				case int:
					e := err.(int)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", e)))
				case int8:
					e := err.(int8)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", e)))
				case int16:
					e := err.(int16)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", e)))
				case int32:
					e := err.(int32)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", e)))
				case int64:
					e := err.(int64)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", e)))
				case *int:
					e := err.(*int)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", *e)))
				case *int8:
					e := err.(*int8)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", *e)))
				case *int16:
					e := err.(*int16)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", *e)))
				case *int32:
					e := err.(*int32)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", *e)))
				case *int64:
					e := err.(*int64)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(fmt.Sprintf("%v", *e)))
				case errorx.BizError:
					e := err.(errorx.BizError)
					c.JSON(http.StatusOK, e.Format())
				case *errorx.BizError:
					e := err.(*errorx.BizError)
					c.JSON(http.StatusOK, e.Format())
				default:
					e := err.(error)
					c.JSON(http.StatusInternalServerError, errorx.NewBizErrorWithMsg(e.Error()))
				}
			}
		}()
		c.Next()
	}
}
