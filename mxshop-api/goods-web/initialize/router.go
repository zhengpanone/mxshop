package initialize

import (
	"github.com/gin-gonic/gin"
	"goods-web/middlewares"
	"goods-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/v1")
	router.InitGoodsRouter(ApiGroup)
	return Router
}
