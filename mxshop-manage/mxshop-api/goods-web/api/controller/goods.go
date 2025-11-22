package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/global"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/request"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/response"
	"go.uber.org/zap"
	"strconv"
)

type GoodsController struct{}

// GetGoodsPageList 获取商品列表
//
//	@Summary		获取商品列表
//	@Description	根据多个查询条件（价格区间、是否热销、是否新品、分类、品牌等）获取商品列表。
//	@Tags			Goods 商品管理
//	@Accept			json
//	@Produce		json
//	@Param			pMin	query		int				false	"最低价格"		default(0)
//	@Param			pMax	query		int				false	"最高价格"		default(0)
//	@Param			ih		query		int				false	"是否热销商品"	default(0)	Enum(0, 1)
//	@Param			ih		query		int				false	"是否新品商品"	default(0)	Enum(0, 1)
//	@Param			ih		query		int				false	"是否Tab商品"	default(0)	Enum(0, 1)
//	@Param			c		query		int				false	"分类ID"		default(0)
//	@Param			page	query		int				false	"页码"		default(1)
//	@Param			size	query		int				false	"每页数量"		default(10)
//	@Param			kw		query		string			false	"搜索关键词"
//	@Param			brand	query		int				false	"品牌ID"	default(0)
//	@Success		200		{object}	utils.Response	"成功获取商品列表"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/goods/getGoodsPageList [get]
func (g *GoodsController) GetGoodsPageList(ctx *gin.Context) {
	goodsPageForm := request.GoodsPageForm{}

	if err := ctx.ShouldBindJSON(&goodsPageForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}

	goodsRequest := &commonpb.GoodsFilterPageRequest{
		PriceMin: goodsPageForm.PriceMin,
		PriceMax: goodsPageForm.PriceMax,
		PageRequest: &commonpb.PageRequest{
			PageNum:  goodsPageForm.PageNum,
			PageSize: goodsPageForm.PageSize,
		},
	}

	// 请求商品service服务
	rsp, err := global.GoodsSrvClient.GoodsPageList(context.Background(), goodsRequest)
	if err != nil {
		zap.S().Errorw("[GetGoodsList]查询【商品列表】失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	goodsList := make([]response.GoodsResponse, 0)
	for _, value := range rsp.List {
		goods := response.GoodsResponse{
			Id:          value.Id,
			Name:        value.Name,
			GoodsBrief:  value.GoodsBrief,
			Description: value.GoodsDesc,
			ShipFree:    value.ShipFree,
			Images:      value.Images,
			DescImages:  value.DescImages,
			FrontImage:  value.GoodsFrontImage,
			ShopPrice:   value.ShopPrice,
			Category: response.CategoryResponse{
				Id:   value.Category.Id,
				Name: value.Category.Name,
			},
			Brand: response.BrandResponse{
				Id:   value.Brand.Id,
				Name: value.Brand.Name,
				Logo: value.Brand.Logo,
			},
			IsHot:  value.IsHot,
			IsNew:  value.IsNew,
			OnSale: value.OnSale,
		}
		goodsList = append(goodsList, goods)
	}
	pageResult := commonUtils.ConvertPage(rsp.Page, goodsList)
	commonResponse.OkWithData(ctx, pageResult)

}

// NewGoods 创建新的商品
//
//	@Summary		创建一个新的商品
//	@Description	根据提交的商品信息创建一个新的商品。
//	@Tags			Goods 商品管理
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"
//	@Param			goods	body		request.GoodsForm	true	"商品信息"
//	@Success		201		{object}	utils.Response	"商品创建成功"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/goods/create [post]
func (*GoodsController) NewGoods(ctx *gin.Context) {
	goodsForm := request.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	goodsClient := global.GoodsSrvClient
	rsp, err := goodsClient.CreateGoods(context.Background(), &commonpb.CreateGoodsRequest{
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		GoodsDesc:       goodsForm.GoodsDesc,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.BrandId,
	})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	// 如何设置库存
	commonResponse.OkWithData(ctx, rsp)
}

// UpdateStatus 更新商品的状态
//
//	@Summary		更新商品的状态（如上架、下架）
//	@Description	根据商品ID更新商品的状态，支持上架、下架等操作。
//	@Tags			Goods 商品管理
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string					true	"认证令牌"
//	@Param			id		path		int						true	"商品ID"
//	@Param			status	body		request.GoodsStatusForm	true	"商品状态信息"
//	@Success		200		{object}	utils.Response			"商品状态更新成功"
//	@Failure		400		{object}	utils.Response			"无效的请求参数"
//	@Failure		500		{object}	utils.Response			"服务器错误"
//	@Router			/v1/goods/status/{id} [put]
func (*GoodsController) UpdateStatus(ctx *gin.Context) {
	goodsStatusForm := request.GoodsStatusForm{}
	if err := ctx.ShouldBindJSON(&goodsStatusForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	id := ctx.Param("id")
	goodsId, err := strconv.ParseUint(id, 10, 64)
	if _, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &commonpb.UpdateGoodsRequest{
		Id:     goodsId,
		IsHot:  *goodsStatusForm.IsHot,
		IsNew:  *goodsStatusForm.IsNew,
		OnSale: *goodsStatusForm.OnSale,
	}); err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	commonResponse.OkWithMsg(ctx, "修改成功")
}
