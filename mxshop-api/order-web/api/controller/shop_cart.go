package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/forms"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/global"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ShopCartApi struct{}

// GetShopCartList 获取购物车列表
//
//	@Summary		获取用户的购物车列表
//	@Description	根据用户的ID获取购物车中的所有商品列表
//	@Tags			ShopCart
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string			true	"认证令牌"	//	用户认证令牌，通常为JWT
//	@Param			user_id	query		int				true	"用户ID"	//	查询参数，用户的ID
//	@Success		200		{object}	utils.Response	"购物车列表获取成功"
//	@Failure		400		{object}	utils.Response	"无效的请求参数"
//	@Failure		404		{object}	utils.Response	"购物车未找到"
//	@Failure		500		{object}	utils.Response	"服务器错误"
//	@Router			/v1/shop-cart/list [get]
func (*ShopCartApi) GetShopCartList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CarItemList(context.Background(), &commonpb.UserInfo{
		Id: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("查询【购物车列表】失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "订单srv")
		return
	}
	ids := make([]int32, 0)
	for _, item := range rsp.Data {
		ids = append(ids, item.GoodsId)
	}
	if len(ids) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}
	// 请求商品服务获取商品信息
	goodsRsp, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &commonpb.BatchGoodsIdInfo{Id: ids})
	if err != nil {
		zap.S().Errorw("批量查询【商品列表】失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	reMap := gin.H{
		"total": rsp.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		for _, goods := range goodsRsp.Data {
			if goods.Id == item.GoodsId {
				tmpMap := map[string]interface{}{}
				tmpMap["id"] = item.Id
				tmpMap["goods_id"] = item.GoodsId
				tmpMap["goods_name"] = goods.Name
				tmpMap["goods_image"] = goods.GoodsFrontImage
				tmpMap["goods_price"] = goods.ShopPrice
				tmpMap["nums"] = item.Nums
				tmpMap["checked"] = item.Checked
				goodsList = append(goodsList, tmpMap)
			}
		}
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

// NewShopCart 添加新商品到购物车
//
//	@Summary		将新商品添加到购物车
//	@Description	根据商品ID和用户ID将商品添加到购物车
//	@Tags			ShopCart
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"	//	用户认证令牌，通常为JWT
//	@Param			user_id		query		int				true	"用户ID"	//	用户ID
//	@Param			product_id	query		int				true	"商品ID"	//	商品ID
//	@Param			quantity	query		int				true	"商品数量"	//	商品数量
//	@Success		200			{object}	utils.Response	"商品添加成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"商品未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/shop-cart/new [post]
func (*ShopCartApi) NewShopCart(ctx *gin.Context) {
	// 添加商品到购物车
	itemForm := forms.ShopCartItemForm{}
	if err := ctx.ShouldBindJSON(&itemForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	// 添加商品到购物车前，检查商品是否存在
	_, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &commonpb.GoodInfoRequest{
		Id: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("查询【商品信息】失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "商品srv")
		return
	}
	// 添加到购物车的数量大于库存
	invRsp, err := global.InventorySrvClient.InvDetail(context.Background(), &commonpb.GoodsInvInfo{
		GoodsId: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("查询【库存信息】失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "库存srv")
		return
	}
	if invRsp.Num < itemForm.Nums {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"nums": "库存不足",
		})
		return
	}
	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CreateCartItem(context.Background(), &commonpb.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: itemForm.GoodsId,
		Nums:    itemForm.Nums,
	})
	if err != nil {
		zap.S().Errorw("添加到购物车失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "订单srv")
		return
	}
	reMap := map[string]interface{}{}
	reMap["id"] = rsp.Id

	commonResponse.OkWithData(ctx, reMap)
}

func (*ShopCartApi) DeleteShopCart(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式错误",
		})
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.OrderSrvClient.DeleteCartItem(context.Background(), &commonpb.CartItemRequest{UserId: int32(userId.(uint)), GoodsId: int32(i)})
	if err != nil {
		zap.S().Errorw("删除购物车失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "订单srv")
		return
	}
	commonResponse.Ok(ctx)
}

func (*ShopCartApi) UpdateShopCart(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式错误",
		})
		return
	}
	itemForm := forms.ShopCartItemUpdateForm{}
	if err := ctx.ShouldBindJSON(&itemForm); err != nil {
		commonUtils.HandleValidatorError(ctx, global.Trans, err)
		return
	}
	userId, _ := ctx.Get("userId")
	request := commonpb.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
		Nums:    itemForm.Nums,
		Checked: false,
	}
	if itemForm.Checked != nil {
		request.Checked = *itemForm.Checked
	}
	_, err = global.OrderSrvClient.UpdateCartItem(context.Background(), &request)
	if err != nil {
		global.Logger.Error("更新购物车失败")
		commonUtils.HandleGrpcErrorToHttp(err, ctx, "订单srv")
		return
	}

	commonResponse.Ok(ctx)
}
