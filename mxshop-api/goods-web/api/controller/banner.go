package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/forms"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/global"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
)

type BannerController struct{}

// ListBanner 获取所有品牌（Banner）列表
//
//	@Summary		获取所有品牌（Banner）列表
//	@Description	返回所有可用的横幅信息列表，支持分页。
//	@Tags			Banner
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"
//	@Param			page	query		int				false	"页码"	default(1)
//	@Param			size	query		int				false	"每页数量"	default(10)
//	@Success		200		{object}	utils.Response	"返回品牌列表"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/banner [get]
func (*BannerController) ListBanner(ctx *gin.Context) {
	rsp, err := global.GoodsSrvClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["index"] = value.Index
		reMap["image"] = value.Image
		reMap["url"] = value.Url
		result = append(result, reMap)
	}
	commonResponse.OkWithData(ctx, result)
}

// NewBanner 创建新的品牌（Banner）
//
//	@Summary		创建新的品牌（Banner）
//	@Description	根据提供的品牌信息创建一个新的品牌（Banner）
//	@Tags			Banner
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string				true	"认证令牌"
//	@Param			banner	body		request.BannerForm	true	"品牌信息"
//	@Success		201		{object}	utils.Response		"品牌创建成功"
//	@Failure		400		{object}	utils.Response		"无效的请求参数"
//	@Failure		500		{object}	utils.Response		"服务器错误"
//	@Router			/v1/banner [post]
func (*BannerController) NewBanner(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}

	rsp, err := global.GoodsSrvClient.CreateBanner(context.Background(),
		&commonpb.BannerRequest{
			Index: int32(bannerForm.Index),
			Url:   bannerForm.Url,
			Image: bannerForm.Image,
		})

	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	response := make(map[string]interface{})
	response["id"] = rsp.Id
	response["index"] = rsp.Index
	response["url"] = rsp.Url
	response["image"] = rsp.Image

	commonResponse.OkWithData(ctx, response)
}

// UpdateBanner 更新轮播图信息
//
//	@Summary		更新轮播图信息
//	@Description	根据轮播图ID更新轮播图的图片URL、链接和描述等信息
//	@Tags			Banner
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string				true	"认证令牌"
//	@Param			id		path		int					true	"轮播图ID"
//	@Param			banner	body		request.BannerForm	true	"更新的轮播图信息"
//	@Success		200		{object}	utils.Response		"轮播图更新成功"
//	@Failure		400		{object}	utils.Response		"无效的请求参数"
//	@Failure		404		{object}	utils.Response		"轮播图未找到"
//	@Failure		500		{object}	utils.Response		"服务器错误"
//	@Router			/v1/banner/{id} [put]
func (*BannerController) UpdateBanner(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.UpdateBanner(context.Background(), &commonpb.BannerRequest{
		Id:    int32(idInt),
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
	})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	commonResponse.Ok(ctx)
}

// DeleteBanner 删除轮播图
//
//	@Summary		删除轮播图
//	@Description	根据轮播图ID删除指定的轮播图
//	@Tags			Banner
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"
//	@Param			id		path		int				true	"轮播图ID"
//	@Success		200		{object}	utils.Response	"轮播图删除成功"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		404		{object}	utils.Response	"轮播图未找到"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/banner/{id} [delete]
func (*BannerController) DeleteBanner(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteBanner(context.Background(), &commonpb.BannerRequest{Id: int32(idInt)})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	commonResponse.Ok(ctx)
}
