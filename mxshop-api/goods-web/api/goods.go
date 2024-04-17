package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goods-web/forms"
	"goods-web/global"
	"goods-web/proto"
	"net/http"
	"strconv"
)

func GetGoodsList(ctx *gin.Context) {
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

	size := ctx.DefaultQuery("page", "10")
	sizeInt, _ := strconv.Atoi(size)
	request.Size = uint32(sizeInt)

	keywords := ctx.DefaultQuery("kw", "")
	request.KeyWords = keywords

	brand := ctx.DefaultQuery("brand", "")
	brandInt, _ := strconv.Atoi(brand)
	request.Brand = int32(brandInt)

	// 请求商品service服务
	r, err := global.GoodsSrvClient.GoodsList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("[GetGoodsList]查询【商品列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": r.Total,
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
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

func NewGoods(ctx *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil {
		HandleValidatorError(ctx, err)
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
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	// 如何设置库存
	ctx.JSON(http.StatusOK, rsp)

}
