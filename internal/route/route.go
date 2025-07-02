// internal/route/route.go
package route

import (
	"github.com/gin-gonic/gin"
	"quanfuxia/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 中间件注册
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger()) // trace_id + zap 日志打印
	r.Use(middleware.I18nMiddleware())
	// 路由分组
	api := r.Group("/api")

	// 用户模块
	RegisterUserRoutes(api)

	return r
}
