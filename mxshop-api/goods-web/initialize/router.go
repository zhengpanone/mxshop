package initialize

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	routers "github.com/zhengpanone/mxshop/goods-web/api/router"
	"github.com/zhengpanone/mxshop/goods-web/global"
	"github.com/zhengpanone/mxshop/goods-web/middleware"
	"net/http"
)

func Routers() *gin.Engine {
	// gin.Default() 已经默认注册了 Logger 和 Recovery
	// 要移除 Gin 默认的 Logger 使用Router.Use()
	//Router := gin.Default()
	// gin.New() 不会自动注册 Logger 和 Recovery
	Router := gin.New()
	// 配置跨域
	Router.Use(
		commonMiddleware.Cors(),
		middleware.Trace(),
		commonMiddleware.CustomLoggerWithConfig([]string{"/health", "/swagger/index.html"}),
		commonMiddleware.GinLogger(global.Logger),
		commonMiddleware.GinRecovery(global.Logger, true),
	)
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
