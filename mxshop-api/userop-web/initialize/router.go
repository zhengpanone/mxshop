package initialize

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	router2 "userop-web/api/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	SwaggerInit(Router)
	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	// 配置跨域
	Router.Use(commonMiddleware.Cors())
	ApiGroup := Router.Group("/v1/userop/")
	router2.InitMessageRouter(ApiGroup)
	router2.InitAddressRouter(ApiGroup)
	router2.InitUserFavRouter(ApiGroup)

	return Router
}
