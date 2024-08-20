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
	Router.Use(middleware.Cors(), middleware.Trace())
	ApiGroup := Router.Group("/v1/goods")
	router.InitGoodsRouter(ApiGroup)
	router.InitCategoryRouter(ApiGroup)
	router.InitBannerRouter(ApiGroup)
	router.InitBrandRouter(ApiGroup)
	return Router
}
