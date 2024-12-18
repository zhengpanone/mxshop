package api

import (
	commonUtils "common/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goods-web/forms"
	"goods-web/global"
	"goods-web/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
)

type CategoryController struct{}

func GetCategoryList(ctx *gin.Context) {
	r, err := global.GoodsSrvClient.GetAllCategoryList(context.Background(), &emptypb.Empty{})

	if err != nil {
		zap.S().Errorw("[GetCategoryList]查询【商品分类列表】失败：", err.Error())
		HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}

	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(r.JsonData), &data)
	if err != nil {
		zap.S().Errorw("[GetCategoryList]查询【商品分类列表】失败：", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"msg": "查询商品分类失败" + err.Error(),
		})
		return
	}
	commonUtils.OkWithData(ctx, data)
}

func Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	reMap := make(map[string]interface{})
	subCategorys := make([]interface{}, 0)
	if r, err := global.GoodsSrvClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: int32(i),
	}); err != nil {
		HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	} else {
		for _, value := range r.SubCategoryList {
			subCategorys = append(subCategorys, map[string]interface{}{
				"id":              value.Id,
				"name":            value.Name,
				"level":           value.Level,
				"parent_category": value.ParentCategory,
				"is_tab":          value.IsTab,
			})
		}
		reMap["id"] = r.Info.Id
		reMap["name"] = r.Info.Name
		reMap["parent_category"] = r.Info.ParentCategory
		reMap["is_tab"] = r.Info.IsTab
		reMap["sub_category_list"] = subCategorys
		ctx.JSON(http.StatusOK, reMap)
	}
	return
}

// CreateCategory
// @Description 创建分类
// @receiver CategoryController
// @Summary 根据给定的参数创建分类
// @Tags Category
// @Accept json
// @Produce json
// @param x-token header string true "x-token header"
// @param forms.CategoryForm body forms.CategoryForm true "category information"
// @Router /v1/goods/category/create [post]
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	// @Success 200 {object} api_helper.Response
	// @Failure 400 {object} api_helper.ErrorResponse
	categoryForm := forms.CategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	goodsClient := global.GoodsSrvClient
	rsp, err := goodsClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:           categoryForm.Name,
		IsTab:          *categoryForm.IsTab,
		Level:          categoryForm.Level,
		ParentCategory: categoryForm.ParentCategory,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	response := make(map[string]interface{})
	response["id"] = rsp.Id
	response["name"] = rsp.Name
	response["parent"] = rsp.ParentCategory
	response["level"] = rsp.Level
	response["is_tab"] = rsp.IsTab

	ctx.JSON(http.StatusOK, response)
}

func UpdateCategory(ctx *gin.Context) {
	categoryForm := forms.UpdateCategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	request := &proto.CategoryInfoRequest{
		Id:   int32(idInt),
		Name: categoryForm.Name,
	}
	if categoryForm.IsTab != nil {
		request.IsTab = *categoryForm.IsTab
	}
	_, err = global.GoodsSrvClient.UpdateCategory(context.Background(), request)
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	ctx.Status(http.StatusOK)
}
