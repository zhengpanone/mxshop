package initialize

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/order-web/api/router"
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
	Router.Use(commonMiddleware.Cors())
	ApiGroup := Router.Group("/v1/order/")
	router.InitOrderRouter(ApiGroup)
	router.InitShopCartRouter(ApiGroup)

	return Router
}
