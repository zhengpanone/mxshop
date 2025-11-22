package router

import (
	"github.com/gin-gonic/gin"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/api/controller"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/global"
)

func InitAddressRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("address").Use(commonMiddleware.JWTAuth(global.ServerConfig.JWTInfo.SigningKey))
	{
		GoodsRouter.GET("/list", controller.GetAddressList)        // 查看地址列表
		GoodsRouter.POST("/create", controller.CreateAddress)      // 新增地址
		GoodsRouter.DELETE("delete/:id", controller.DeleteAddress) // 删除地址
		GoodsRouter.PUT("update/:id", controller.UpdateAddress)    // 修改地址
	}

}
