package router

import (
	"github.com/gin-gonic/gin"
	"oss-web/api"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("oss")
	{
		OssRouter.GET("token", api.GenerateUploadToken)
		OssRouter.POST("callback", api.HandlerRequest)
	}
}
