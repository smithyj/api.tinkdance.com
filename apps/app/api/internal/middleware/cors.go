package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
	"tinkdance/apps/app/api/internal/svc"
	"tinkdance/pkg/tracex"
)

var allowOrigins = []string{
	"tinkdance.com",
}

func Cors(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	cfg := cors.DefaultConfig()

	// 允许携带 Cookie
	cfg.AllowCredentials = true

	// 允许请求方法
	cfg.AddAllowMethods("GET", "POST", "PUT", "DELETE", "PATCH")

	// 允许设置 Header 头
	cfg.AddAllowHeaders(tracex.TraceKey)

	// 允许访问 Header 头
	cfg.AddExposeHeaders(tracex.TraceKey, tracex.RequestKey)

	// 跨域访问认证
	cfg.AllowOriginFunc = func(origin string) bool {
		if svcCtx.Config.Mode == "debug" {
			return true
		}
		p, err := url.Parse(origin)
		if err != nil {
			return false
		}
		for _, v := range allowOrigins {
			if strings.HasSuffix(p.Host, v) {
				return true
			}
		}
		return false
	}

	return cors.New(cfg)
}
