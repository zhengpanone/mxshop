package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/mxshop/mxshop-api/oss-web/api/controller"
)

func InitOssRouter(Router *gin.RouterGroup) {
	OssRouter := Router.Group("oss")
	{
		OssRouter.GET("token", controller.GenerateUploadToken)
		OssRouter.POST("callback", controller.HandlerRequest)
		OssRouter.POST("upload", controller.UploadFile)
	}
}
