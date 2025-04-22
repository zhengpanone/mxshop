package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/mxshop/userop-web/forms"
	"github.com/zhengpanone/mxshop/userop-web/global"
	"github.com/zhengpanone/mxshop/userop-web/proto"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// GetFavList 获取用户收藏列表
//
//	@Summary		获取用户收藏列表
//	@Description	根据用户收藏列表。
//	@Tags			UserFav
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/userfavs/list [get]
func GetFavList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	userFavRsp, err := global.UserFavSrvClient.GetFavList(context.Background(), &proto.UserFavRequest{UserId: int32(userId.(uint))})
	if err != nil {
		zap.S().Errorw("获取收藏列表失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}
	goodsIds := make([]int32, 0)
	for _, item := range userFavRsp.Data {
		goodsIds = append(goodsIds, item.GoodsId)
	}
	if len(goodsIds) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}
	goods, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: goodsIds,
	})
	if err != nil {
		zap.S().Errorw("批量查询商品列表失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}
	reMap := map[string]interface{}{
		"total": userFavRsp.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, item := range userFavRsp.Data {
		data := gin.H{
			"id": item.GoodsId,
		}
		for _, good := range goods.Data {
			for item.GoodsId == good.Id {
				data["name"] = good.Name
				data["shopPrice"] = good.ShopPrice
			}
		}
		goodsList = append(goodsList, data)
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

// AddUserFav 		新增收藏
//
//	@Summary		新增用户收藏
//	@Description	新增用户收藏。
//	@Tags			UserFav
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/userfavs/create [post]
func AddUserFav(ctx *gin.Context) {
	userFavForm := forms.UserFavForm{}
	if err := ctx.ShouldBindJSON(&userFavForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err := global.UserFavSrvClient.AddUserFav(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: userFavForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("新增收藏列表失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "添加到收藏夹成功",
	})

}

// DeleteUserFav 		删除收藏
//
//	@Summary		新增用户收藏
//	@Description	新增用户收藏。
//	@Tags			UserFav
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/userfavs/delete [delete]
func DeleteUserFav(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, _ := ctx.Get("userId")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.UserFavSrvClient.DeleteUserFav(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
	})
	if err != nil {
		zap.S().Errorw("删除收藏列表失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "取消收藏成功",
	})
}

// GetUserFavDetail 		获取用户收藏详情
//
//	@Summary		获取用户收藏详情
//	@Description	获取用户收藏详情。
//	@Tags			UserFav
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/userfavs/detail [get]
func GetUserFavDetail(ctx *gin.Context) {
	goodsId := ctx.Param("id")
	userId, _ := ctx.Get("userId")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.UserFavSrvClient.GetUserFavDetail(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(goodsIdInt),
	})
	if err != nil {
		zap.S().Errorw("获取商品收藏状态失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}

	ctx.Status(http.StatusOK)
}
