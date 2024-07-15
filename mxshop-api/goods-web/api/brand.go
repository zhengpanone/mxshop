package api

import (
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
		HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})

	if err != nil {
		HandleGrpcErrorToHttp(err, ctx, "商品srv")
	}
	response := make(map[string]interface{})
	response["id"] = rsp.Id
	response["name"] = rsp.Name
	response["logo"] = rsp.Logo
	ctx.JSON(http.StatusOK, response)
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
		HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	ctx.Status(http.StatusOK)
}

func UpdateBrand(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(brandForm); err != nil {
		HandleValidatorError(ctx, err)
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
		HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	ctx.Status(http.StatusOK)
}

func ListBrand(ctx *gin.Context) {
	page := ctx.DefaultQuery("pn", "1")
	pageInt, _ := strconv.Atoi(page)
	size := ctx.DefaultQuery("size", "10")
	sizeInt, _ := strconv.Atoi(size)

	rsp, err := global.GoodsSrvClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Page: int32(pageInt),
		Size: int32(sizeInt),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	result := make([]interface{}, 0)
	response := make(map[string]interface{})
	response["total"] = rsp.Total
	for _, value := range rsp.Data[pageInt : pageInt*sizeInt+sizeInt] {
		responseMap := make(map[string]interface{})
		responseMap["id"] = value.Id
		responseMap["name"] = value.Name
		responseMap["logo"] = value.Logo
		result = append(result, responseMap)
	}
	response["data"] = result
	ctx.JSON(http.StatusOK, result)

}
