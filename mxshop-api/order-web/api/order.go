package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"net/http"
	"order-web/forms"
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
	orderForm := forms.CreateOrderForm{}
	if err := ctx.ShouldBindJSON(&orderForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	userId, _ := ctx.Get("userId")
	request := proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Name:    orderForm.Name,
		Address: orderForm.Address,
		Post:    orderForm.Post,
		Mobile:  orderForm.Mobile,
	}
	rsp, err := global.OrderSrvClient.CreateOrder(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("新建订单详情失败")
		HandleGrpcErrorToHttp(err, ctx, "订单srv")
		return
	}
	//  返回支付宝的支付url
	client, err := alipay.New(global.ServerConfig.AliPayConfig.AppId, global.ServerConfig.AliPayConfig.PrivateKey, false)
	if err != nil {
		zap.S().Errorw("实例化支付宝客户端支付失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err = client.LoadAliPayPublicKey(global.ServerConfig.AliPayConfig.AliPublicKey)
	if err != nil {
		zap.S().Errorw("加载支付宝公钥失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = global.ServerConfig.AliPayConfig.NotifyUrl
	p.ReturnURL = global.ServerConfig.AliPayConfig.ReturnUrl
	p.Subject = "慕学生鲜订单-" + rsp.OrderSn
	p.OutTradeNo = rsp.OrderSn
	p.TotalAmount = strconv.FormatFloat(float64(rsp.Total), 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("生成支付url失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         rsp.Id,
		"alipay_url": url.String(),
	})

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
	//  返回支付宝的支付url
	client, err := alipay.New(global.ServerConfig.AliPayConfig.AppId, global.ServerConfig.AliPayConfig.PrivateKey, false)
	if err != nil {
		zap.S().Errorw("实例化支付宝客户端支付失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err = client.LoadAliPayPublicKey(global.ServerConfig.AliPayConfig.AliPublicKey)
	if err != nil {
		zap.S().Errorw("加载支付宝公钥失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = global.ServerConfig.AliPayConfig.NotifyUrl
	p.ReturnURL = global.ServerConfig.AliPayConfig.ReturnUrl
	p.Subject = "慕学生鲜订单-" + rsp.OrderInfo.OrderSn
	p.OutTradeNo = rsp.OrderInfo.OrderSn
	p.TotalAmount = strconv.FormatFloat(float64(rsp.OrderInfo.Total), 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("生成支付url失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	reMap["alipay_url"] = url.String()
	ctx.JSON(http.StatusOK, reMap)
}
