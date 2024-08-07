package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"userop-web/middlewares"
	"userop-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	// 配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/v1/userop/")
	router.InitMessageRouter(ApiGroup)
	router.InitAddressRouter(ApiGroup)
	router.InitUserFavRouter(ApiGroup)

	return Router
}
