package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/handler"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("oss")
	{
		OssRouter.GET("token", handler.Token)
		OssRouter.POST("callback", handler.HandlerRequest)
	}
}
