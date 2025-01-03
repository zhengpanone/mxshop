package controller

import (
	commonUtils "common/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goods-web/forms"
	"goods-web/global"
	"goods-web/proto"
	"math"
	"strconv"
)

type GoodsController struct{}

// GetGoodsList 获取商品列表
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
//	@Router			/v1/goods/list [get]
func (*GoodsController) GetGoodsList(ctx *gin.Context) {
	request := &proto.GoodsFilterRequest{}
	priceMin := ctx.DefaultQuery("pMin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	request.PriceMin = int32(priceMinInt)

	priceMax := ctx.DefaultQuery("pMax", "0")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	request.PriceMax = int32(priceMaxInt)

	isHot := ctx.DefaultQuery("ih", "0")
	if isHot == "1" {
		request.IsHot = true
	}

	isNew := ctx.DefaultQuery("ih", "0")
	if isNew == "1" {
		request.IsNew = true
	}
	isTab := ctx.DefaultQuery("ih", "0")
	if isTab == "1" {
		request.IsTab = true
	}

	categoryId := ctx.DefaultQuery("c", "0")
	categoryIdInt, _ := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	request.Page = uint32(pageInt)

	size := ctx.DefaultQuery("size", "10")
	sizeInt, _ := strconv.Atoi(size)
	request.Size = uint32(sizeInt)

	keywords := ctx.DefaultQuery("kw", "")
	request.KeyWords = keywords

	brand := ctx.DefaultQuery("brand", "")
	brandInt, _ := strconv.Atoi(brand)
	request.Brand = int32(brandInt)

	// 请求商品service服务
	r, err := global.GoodsSrvClient.GoodsList(context.WithValue(context.Background(), "ginContext", ctx), request)
	if err != nil {
		zap.S().Errorw("[GetGoodsList]查询【商品列表】失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}

	reMap := map[string]interface{}{
		"total":     r.Total,
		"totalPage": int(math.Ceil(float64(r.Total) / float64(request.Size))),
		"pageNum":   request.Page,
		"pageSize":  request.Size,
	}
	goodsList := make([]interface{}, 0)
	for _, value := range r.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}
	reMap["list"] = goodsList
	commonUtils.OkWithData(ctx, reMap)

}

// NewGoods 创建新的商品
//
//	@Summary		创建一个新的商品
//	@Description	根据提交的商品信息创建一个新的商品。
//	@Tags			Goods 商品管理
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"
//	@Param			goods	body		forms.GoodsForm	true	"商品信息"
//	@Success		201		{object}	utils.Response	"商品创建成功"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/goods/create [post]
func (*GoodsController) NewGoods(ctx *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	goodsClient := global.GoodsSrvClient
	rsp, err := goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
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
		BrandId:         goodsForm.Brand,
	})
	if err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	// 如何设置库存
	commonUtils.OkWithData(ctx, rsp)
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
//	@Param			status	body		forms.GoodsStatusForm	true	"商品状态信息"
//	@Success		200		{object}	utils.Response			"商品状态更新成功"
//	@Failure		400		{object}	utils.Response			"无效的请求参数"
//	@Failure		500		{object}	utils.Response			"服务器错误"
//	@Router			/v1/goods/status/{id} [put]
func (*GoodsController) UpdateStatus(ctx *gin.Context) {
	goodsStatusForm := forms.GoodsStatusForm{}
	if err := ctx.ShouldBindJSON(&goodsStatusForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if _, err = global.GoodsSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:     int32(i),
		IsHot:  *goodsStatusForm.IsHot,
		IsNew:  *goodsStatusForm.IsNew,
		OnSale: *goodsStatusForm.OnSale,
	}); err != nil {
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	commonUtils.OkWithMsg(ctx, "修改成功")
}
