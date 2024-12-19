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

// GetCategoryList 获取分类列表
// @Summary 获取分类列表
// @Description 根据条件获取分类列表，支持分页和过滤
// @receiver CategoryController
// @Tags Category 分类管理
// @Accept json
// @Produce json
// @Param x-token header string true "认证令牌"
// @Param page query int false "页码，默认为1" default(1)
// @Param page_size query int false "每页数量，默认为10" default(10)
// @Param name query string false "按名称过滤分类"
// @Success 200 {object} utils.Response "成功获取分类列表的响应数据"
// @Failure 400 {object} utils.Response "无效的请求参数"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 500 {object} utils.Response "服务器错误"
// @Router /v1/goods/category/list [get]
func (c *CategoryController) GetCategoryList(ctx *gin.Context) {
	r, err := global.GoodsSrvClient.GetAllCategoryList(context.Background(), &emptypb.Empty{})

	if err != nil {
		zap.S().Errorw("[GetCategoryList]查询【商品分类列表】失败：", err.Error())
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
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

// Detail 获取分类详情
// @Summary 获取指定分类的详细信息
// @Description 根据分类ID获取指定分类的详细信息，包括名称、父级分类等
// @Tags Category 分类管理
// @Accept json
// @Produce json
// @Param x-token header string true "认证令牌"
// @Param id path int true "分类ID"
// @Success 200 {object} utils.Response "成功获取分类详情"
// @Failure 400 {object} utils.Response "无效的请求参数"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 404 {object} utils.Response "分类未找到"
// @Failure 500 {object} utils.Response "服务器错误"
// @Router /v1/goods/category/detail/{id} [get]
func (c *CategoryController) Detail(ctx *gin.Context) {
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
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
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
		commonUtils.OkWithData(ctx, reMap)
	}
	return
}

// CreateCategory 创建分类
// @Summary 根据给定的参数创建分类
// @Description 创建分类
// @receiver CategoryController
// @Tags Category 分类管理
// @Accept json
// @Produce json
// @param x-token header string true "x-token header"
// @param forms.CategoryForm body forms.CategoryForm true "分类信息"
// @Success 201 {object} utils.Response "创建成功的响应数据"
// @Failure 400 {object} utils.Response "无效的请求参数"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 500 {object} utils.Response "服务器错误"
// @Router /v1/goods/category/create [post]
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	// @Success 200 {object} api_helper.Response
	// @Failure 400 {object} api_helper.ErrorResponse
	categoryForm := forms.CategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
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
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	response := make(map[string]interface{})
	response["id"] = rsp.Id
	response["name"] = rsp.Name
	response["parent"] = rsp.ParentCategory
	response["level"] = rsp.Level
	response["is_tab"] = rsp.IsTab

	commonUtils.OkWithData(ctx, response)
}

// UpdateCategory 更新分类信息
// @Summary 更新指定分类的信息
// @Description 根据分类ID更新分类信息，支持修改名称、父级分类等属性
// @receiver CategoryController
// @Tags Category 分类管理
// @Accept json
// @Produce json
// @Param x-token header string true "认证令牌"
// @Param id path int true "分类ID"
// @Param category body forms.CategoryForm true "分类信息"
// @Success 200 {object} utils.Response "分类信息更新成功"
// @Failure 400 {object} utils.Response "无效的请求参数"
// @Failure 401 {object} utils.Response "未授权"
// @Failure 404 {object} utils.Response "分类未找到"
// @Failure 500 {object} utils.Response "服务器错误"
// @Router /v1/goods/category/update/{id} [post]
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	categoryForm := forms.UpdateCategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
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
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	commonUtils.Ok(ctx)
}
