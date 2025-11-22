package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/api/controller"
)

func InitDictRouter(router *gin.RouterGroup) {
	DictRouter := router.Group("dict")
	dictController := &controller.DictTypeController{}
	{
		DictRouter.POST("createDictType", dictController.CreateDictType)

	}
}
