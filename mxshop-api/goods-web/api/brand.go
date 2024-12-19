package api

import (
	commonUtils "common/utils"
	"context"
	"github.com/gin-gonic/gin"
	"goods-web/forms"
	"goods-web/global"
	"goods-web/proto"
	"net/http"
	"strconv"
)

func NewBrand(ctx *gin.Context) {
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

func DeleteBrand(ctx *gin.Context) {
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

func UpdateBrand(ctx *gin.Context) {
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

func ListBrand(ctx *gin.Context) {
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
	response["data"] = result
	commonUtils.OkWithData(ctx, response)
}

func GetCategoryBrandList(ctx *gin.Context) {
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

func UpdateCategoryBrand(ctx *gin.Context) {

}

func CategoryBrandList(ctx *gin.Context) {
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

func NewCategoryBrand(ctx *gin.Context) {

}

func DeleteCategoryBrand(ctx *gin.Context) {

}
