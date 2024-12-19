package initialize

import (
	"common/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	router2 "user-web/api/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	SwaggerInit(Router)
	Router.Use(middleware.Recovery())
	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	// 配置跨域
	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("/v1")
	router2.InitBaseRouter(ApiGroup)
	router2.InitUserRouter(ApiGroup)
	return Router
}
