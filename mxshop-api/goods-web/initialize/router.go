package initialize

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	routers "github.com/zhengpanone/mxshop/goods-web/api/router"
	"github.com/zhengpanone/mxshop/goods-web/middleware"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 配置跨域
	Router.Use(gin.LoggerWithConfig(gin.LoggerConfig{}), commonMiddleware.Cors(), middleware.Trace())
	SwaggerInit(Router)
	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	ApiGroup := Router.Group("/v1/goods")
	routers.InitGoodsRouter(ApiGroup)
	routers.InitCategoryRouter(ApiGroup)
	routers.InitBannerRouter(ApiGroup)
	routers.InitBrandRouter(ApiGroup)
	return Router
}
