package initialize

import (
	"github.com/gin-gonic/gin"
	"goods-web/middleware"
	"goods-web/router"
	"net/http"
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
	Router.Use(middleware.Cors())
	ApiGroup := Router.Group("/v1")
	router.InitOssRouter(ApiGroup)
	return Router
}
