package initialize

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	"github.com/zhengpanone/mxshop/mxshop-api/oss-web/api/router"
	"net/http"
)

func Routers() *gin.Engine {
	// gin.Default() 已经默认注册了 Logger 和 Recovery
	// 要移除 Gin 默认的 Logger 使用Router.Use()
	//Router := gin.Default()
	// gin.New() 不会自动注册 Logger 和 Recovery
	Router := gin.New()
	
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
