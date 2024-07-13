package router

import (
	"github.com/gin-gonic/gin"
	"user-web/api"
)

func InitBaseRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("base")
	{
		BaseRouter.GET("ping", api.Ping)
		BaseRouter.GET("captcha", api.GenerateCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)
	}
}
