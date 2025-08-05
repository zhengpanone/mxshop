package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/api/controller"
)

func InitBaseRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("base")
	{
		BaseRouter.GET("ping", controller.Ping)
		BaseRouter.GET("captcha", controller.GenerateCaptcha)
		BaseRouter.POST("send_sms", controller.SendSms)
	}
}
