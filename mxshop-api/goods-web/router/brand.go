package router

import (
	"github.com/gin-gonic/gin"
	"goods-web/api"
	"goods-web/middleware"
)

func InitBrandRouter(router *gin.RouterGroup) {
	BrandRouter := router.Group("brand")
	{
		BrandRouter.GET("list", api.ListBrand)
		BrandRouter.POST("create", middleware.JWTAuth(), middleware.IsAdmin(), api.NewBrand)
		BrandRouter.PUT("update", middleware.JWTAuth(), middleware.IsAdmin(), api.UpdateBrand)
		BrandRouter.DELETE("delete", middleware.JWTAuth(), middleware.IsAdmin(), api.DeleteBrand)
	}
	CategoryBrandRouter := router.Group("category-brand")
	{
		CategoryBrandRouter.GET("list", api.CategoryBrandList)     // 类别品牌列表页
		CategoryBrandRouter.DELETE(":id", api.DeleteCategoryBrand) // 删除类别品牌
		CategoryBrandRouter.POST("create", api.NewCategoryBrand)   //新建类别品牌
		CategoryBrandRouter.PUT(":id", api.UpdateCategoryBrand)    //修改类别品牌
		CategoryBrandRouter.GET(":id", api.GetCategoryBrandList)   //获取分类的品牌
	}
}
