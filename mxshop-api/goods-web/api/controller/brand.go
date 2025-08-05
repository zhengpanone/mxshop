package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/forms"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/global"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/proto"
	"net/http"
	"strconv"
)

type BrandController struct{}

// NewBrand 创建新的品牌
//
//	@Summary		创建一个新的品牌
//	@Description	根据提交的品牌信息创建一个新的品牌。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"
//	@Param			brand	body		forms.BrandForm	true	"品牌信息"
//	@Success		201		{object}	utils.Response	"品牌创建成功"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/create [post]
func (*BrandController) NewBrand(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(&brandForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})

	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
	}
	response := make(map[string]interface{})
	response["id"] = rsp.Id
	response["name"] = rsp.Name
	response["logo"] = rsp.Logo
	commonUtils.OkWithData(ctx, response)
}

// DeleteBrand 删除指定的品牌
//
//	@Summary		删除指定的品牌
//	@Description	根据品牌ID删除指定的品牌。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"
//	@Param			id		path		int				true	"品牌ID"
//	@Success		200		{object}	utils.Response	"品牌删除成功"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		404		{object}	utils.Response	"品牌未找到"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/delete/{id} [delete]
func (*BrandController) DeleteBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteBrand(context.Background(), &proto.BrandRequest{Id: int32(idInt)})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	commonUtils.Ok(ctx)
}

// UpdateBrand 更新品牌信息
//
//	@Summary		更新指定品牌的信息
//	@Description	根据品牌ID更新指定品牌的名称和Logo等信息。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"
//	@Param			id		path		int				true	"品牌ID"
//	@Param			brand	body		forms.BrandForm	true	"更新品牌信息"
//	@Success		200		{object}	utils.Response	"品牌更新成功"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		404		{object}	utils.Response	"品牌未找到"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/update/{id} [put]
func (*BrandController) UpdateBrand(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(brandForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteBrand(context.Background(), &proto.BrandRequest{
		Id:   int32(idInt),
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	commonUtils.Ok(ctx)
}

// ListBrand 获取品牌列表
//
//	@Summary		获取所有品牌的列表
//	@Description	获取品牌列表，可以根据分页参数进行分页。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"
//	@Param			page	query		int				false	"页码"	default(1)
//	@Param			size	query		int				false	"每页数量"	default(10)
//	@Success		200		{object}	utils.Response	"品牌列表获取成功"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/list [get]
func (*BrandController) ListBrand(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	size := ctx.DefaultQuery("size", "10")
	sizeInt, _ := strconv.Atoi(size)

	rsp, err := global.GoodsSrvClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Page: int32(pageInt),
		Size: int32(sizeInt),
	})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	result := make([]interface{}, 0)
	response := make(map[string]interface{})
	response["total"] = rsp.Total
	for _, value := range rsp.Data {
		responseMap := make(map[string]interface{})
		responseMap["id"] = value.Id
		responseMap["name"] = value.Name
		responseMap["logo"] = value.Logo
		result = append(result, responseMap)
	}
	response["list"] = result
	commonUtils.OkWithData(ctx, response)
}

// GetCategoryBrandList 获取指定分类下的品牌列表
//
//	@Summary		获取指定分类下的品牌列表
//	@Description	根据分类ID获取该分类下的所有品牌信息。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"品牌列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/category/{categoryId} [get]
func (*BrandController) GetCategoryBrandList(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	resp, err := global.GoodsSrvClient.GetCategoryBrandList(context.Background(), &proto.CategoryBrandInfoRequest{
		Id: int32(categoryId),
	})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	result := make([]interface{}, 0)
	for _, value := range resp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["name"] = value.Name
		reMap["logo"] = value.Logo

		result = append(result, reMap)
	}
	commonUtils.OkWithData(ctx, result)
}

// UpdateCategoryBrand 更新分类品牌信息
//
//	@Summary		更新分类的品牌信息
//	@Description	根据给定的分类ID和品牌ID更新分类下的品牌信息。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Param			brandId		path		int				true	"品牌ID"
//	@Success		200			{object}	utils.Response	"更新成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类或品牌未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/category/{categoryId}/brand/{brandId} [put]
func (*BrandController) UpdateCategoryBrand(ctx *gin.Context) {

}

// CategoryBrandList 获取指定分类下的品牌列表
//
//	@Summary		获取指定分类下的所有品牌信息
//	@Description	根据分类ID获取该分类下所有的品牌信息。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"品牌列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/category/{categoryId}/brands [get]
func (*BrandController) CategoryBrandList(ctx *gin.Context) {
	//所有的list返回的数据结构
	/*
		{
			"total": 100,
			"data":[{},{}]
		}
	*/
	rsp, err := global.GoodsSrvClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	reMap := map[string]interface{}{
		"total": rsp.Total,
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["category"] = map[string]interface{}{
			"id":   value.Category.Id,
			"name": value.Category.Name,
		}
		reMap["brand"] = map[string]interface{}{
			"id":   value.Brand.Id,
			"name": value.Brand.Name,
			"logo": value.Brand.Logo,
		}

		result = append(result, reMap)
	}

	reMap["data"] = result
	commonUtils.OkWithData(ctx, reMap)
}

// NewCategoryBrand 为分类添加品牌
//
//	@Summary		为指定分类添加新的品牌
//	@Description	根据分类ID为该分类添加新的品牌信息。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"品牌添加成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/category/{categoryId}/brand [post]
func (*BrandController) NewCategoryBrand(ctx *gin.Context) {

}

// DeleteCategoryBrand 删除指定分类下的品牌
//
//	@Summary		删除指定分类下的品牌
//	@Description	根据分类ID和品牌ID删除该分类下的品牌。
//	@Tags			Brand
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Param			brandId		path		int				true	"品牌ID"
//	@Success		200			{object}	utils.Response	"品牌删除成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类或品牌未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/brand/category/{categoryId}/brand/{brandId} [delete]
func (*BrandController) DeleteCategoryBrand(ctx *gin.Context) {

}
