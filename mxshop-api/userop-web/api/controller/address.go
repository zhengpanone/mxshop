package controller

import (
	customClaims "common/claims"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"userop-web/forms"
	"userop-web/global"
	"userop-web/proto"
)

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
