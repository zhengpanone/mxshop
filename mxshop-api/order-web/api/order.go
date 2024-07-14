package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"order-web/global"
	"order-web/models"
	"order-web/proto"
	"strconv"
)

func GetOrderList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	claims, _ := ctx.Get("claims")
	// 管理员查询所有订单
	model := claims.(*models.CustomClaims)
	request := proto.OrderFilterRequest{}
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	request.Page = int32(pageInt)

	size := ctx.DefaultQuery("size", "10")
	sizeInt, _ := strconv.Atoi(size)
	request.Size = int32(sizeInt)

	rsp, err := global.OrderSrvClient.OrderList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单列表失败")
		HandleGrpcErrorToHttp(err, ctx, "订单srv")
		return
	}
	reMap := map[string]interface{}{
		"total": rsp.Total,
	}
	orderList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		orderList = append(orderList, map[string]interface{}{
			"id":      item.Id,
			"status":  item.Status,
			"user":    item.UserId,
			"post":    item.Post,
			"payType": item.PayType,
			"orderSn": item.OrderSn,
			"addTime": item.AddTime,
			"total":   item.Total,
			"address": item.Address,
			"mobile":  item.Mobile,
			"name":    item.Name,
		})
	}
	reMap["data"] = orderList
	ctx.JSON(http.StatusOK, reMap)
}

func NewOrder(ctx *gin.Context) {
	/*goodsForm := forms.GoodsForm{}
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
	})*/

	/*if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	// 如何设置库存
	ctx.JSON(http.StatusOK, rsp)*/

}

func GetOrderDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, _ := ctx.Get("userId")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式错误",
		})
		return
	}
	claims, _ := ctx.Get("claims")
	// 管理员查询所有订单
	model := claims.(*models.CustomClaims)
	request := proto.OrderRequest{
		Id: int32(i),
	}
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	rsp, err := global.OrderSrvClient.OrderDetail(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单详情失败")
		HandleGrpcErrorToHttp(err, ctx, "订单srv")
		return
	}

	reMap := gin.H{}
	reMap["id"] = rsp.OrderInfo.Id
	reMap["status"] = rsp.OrderInfo.Status
	reMap["userId"] = rsp.OrderInfo.UserId
	reMap["post"] = rsp.OrderInfo.Post
	reMap["total"] = rsp.OrderInfo.Total
	reMap["address"] = rsp.OrderInfo.Address
	reMap["name"] = rsp.OrderInfo.Name
	reMap["mobile"] = rsp.OrderInfo.Mobile
	reMap["payType"] = rsp.OrderInfo.PayType
	reMap["orderSn"] = rsp.OrderInfo.OrderSn
	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Goods {
		tmpMap := gin.H{
			"id":    item.GoodsId,
			"image": item.GoodsImage,
			"price": item.GoodsPrice,
			"nums":  item.Nums,
		}
		goodsList = append(goodsList, tmpMap)
	}
	reMap["goods"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}
