package initialize

import (
	"github.com/gin-gonic/gin"
	middleware2 "goods-web/api/middleware"
	router2 "goods-web/api/router"
	"net/http"
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
	Router.Use(middleware2.Cors(), middleware2.Trace())
	ApiGroup := Router.Group("/v1/goods")
	router2.InitGoodsRouter(ApiGroup)
	router2.InitCategoryRouter(ApiGroup)
	router2.InitBannerRouter(ApiGroup)
	router2.InitBrandRouter(ApiGroup)
	return Router
}
