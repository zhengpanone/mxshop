package initialize

import (
	"github.com/gin-gonic/gin"
	"goods-web/middlewares"
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
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/v1")
	router.InitGoodsRouter(ApiGroup)
	return Router
}
