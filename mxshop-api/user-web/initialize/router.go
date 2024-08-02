package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-web/middlewares"
	"user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	SwaggerInit(Router)
	Router.Use(middlewares.Recovery())
	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	// 配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/v1")
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return Router
}
