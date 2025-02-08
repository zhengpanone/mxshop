package initialize

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/common/middleware"
	"github.com/zhengpanone/mxshop/user-web/api/router"
	"github.com/zhengpanone/mxshop/user-web/middleware"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	SwaggerInit(Router)
	Router.Use(middleware.Recovery())
	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	// 配置跨域
	Router.Use(commonMiddleware.Cors())
	ApiGroup := Router.Group("/v1")
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return Router
}
