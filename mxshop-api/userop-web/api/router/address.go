package router

import (
	commonMiddleware "common/middleware"
	"github.com/gin-gonic/gin"
	"userop-web/api/controller"
	"userop-web/global"
)

func InitAddressRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("address").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", controller.GetAddressList)   // 查看地址
		GoodsRouter.POST("/create", controller.CreateAddress) // 新增地址
		GoodsRouter.DELETE("/:id", controller.DeleteAddress)  // 删除地址
		GoodsRouter.PUT("/:id", controller.UpdateAddress)     // 修改地址
	}

}
