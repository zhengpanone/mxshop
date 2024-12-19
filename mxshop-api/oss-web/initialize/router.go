package initialize

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"oss-web/api/router"
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
	Router.Use(commonMiddleware.Cors())
	ApiGroup := Router.Group("/v1")
	router.InitOssRouter(ApiGroup)
	return Router
}
