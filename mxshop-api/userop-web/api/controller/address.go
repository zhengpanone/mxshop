package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	customClaims "github.com/zhengpanone/mxshop/mxshop-api/common/claims"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/forms"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/global"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/proto"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// GetAddressList 获取用户地址列表
//
//	@Summary		获取用户地址列表
//	@Description	根据用户ID获取该用户的地址信息。
//	@Tags			Address
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/address/list [get]
func GetAddressList(ctx *gin.Context) {
	request := proto.AddressRequest{}
	userId, _ := ctx.Get("userId")
	claims, _ := ctx.Get("claims")
	// 管理员查询所有订单
	model := claims.(*customClaims.CustomClaims)
	if model.AuthorityId != 2 {
		request.UserId = int32(userId.(uint))
	}
	rsp, err := global.AddressSrvClient.GetAddressList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取地址列表失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}
	reMap := map[string]interface{}{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		item := make(map[string]interface{})
		item["id"] = value.Id
		item["userId"] = value.UserId
		item["province"] = value.Province
		item["city"] = value.City
		item["district"] = value.District
		item["address"] = value.Address
		item["singerName"] = value.SingerName
		item["singerMobile"] = value.SingerMobile
		item["id"] = value.Id
		result = append(result, item)
	}
	reMap["data"] = result
	ctx.JSON(http.StatusOK, reMap)
}

// CreateAddress 新增用户地址
//
//	@Summary		新增用户地址
//	@Description	创建用户地址信息
//	@Tags			Address
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/address/create [post]
func CreateAddress(ctx *gin.Context) {
	addressForm := forms.AddressForm{}
	if err := ctx.ShouldBindJSON(addressForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.AddressSrvClient.CreateAddress(context.Background(), &proto.AddressRequest{
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SingerMobile: addressForm.SignerMobile,
		SingerName:   addressForm.SignerName,
	})
	if err != nil {
		zap.S().Errorw("新建地址失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}
	reMap := make(map[string]interface{})
	reMap["id"] = rsp.Id
	ctx.JSON(http.StatusOK, reMap)
}

// DeleteAddress 删除用户地址
//
//	@Summary		删除用户地址
//	@Description	删除用户地址信息
//	@Tags			Address
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/address/delete [delete]
func DeleteAddress(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.AddressSrvClient.DeleteAddress(context.Background(), &proto.AddressRequest{Id: int32(i)})
	if err != nil {
		zap.S().Errorw("删除地址失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除地址成功",
	})
}

// UpdateAddress 更新用户地址
//
//	@Summary		更新用户地址
//	@Description	更新用户地址信息
//	@Tags			Address
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/address/update [put]
func UpdateAddress(ctx *gin.Context) {
	addressForm := forms.AddressForm{}
	if err := ctx.ShouldBindJSON(addressForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.AddressSrvClient.UpdateAddress(context.Background(), &proto.AddressRequest{
		Id:           int32(i),
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SingerMobile: addressForm.SignerMobile,
		SingerName:   addressForm.SignerName,
	})
	if err != nil {
		zap.S().Errorw("更新地址失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新地址成功",
	})
}
