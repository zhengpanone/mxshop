package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/common/middleware"
	"net/http"
	"user-web/router"
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
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return Router
}
