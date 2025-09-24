package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	customClaims "github.com/zhengpanone/mxshop/mxshop-api/common/claims"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/forms"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/global"
	"go.uber.org/zap"
	"net/http"
)

// GetMessageList 获取用户地址列表
//
//	@Summary		获取留言列表
//	@Description	获取留言列表。
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	commonUtils.Response{}	"地址列表获取成功"
//	@Failure		400			{object}	commonUtils.Response	"无效的请求参数"
//	@Failure		404			{object}	commonUtils.Response	"分类未找到"
//	@Failure		500			{object}	commonUtils.Response	"服务器错误"
//	@Router			/v1/userop/message/list [get]
func GetMessageList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	claims, _ := ctx.Get("claims")
	// 管理员查询所有订单
	model := claims.(*customClaims.CustomClaims)
	request := commonpb.MessageRequest{}
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}
	rsp, err := global.MessageSrvClient.MessageList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取留言失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}
	reMap := map[string]interface{}{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["userId"] = value.UserId
		reMap["messageType"] = value.MessageType
		reMap["subject"] = value.Subject
		reMap["message"] = value.Message
		reMap["file"] = value.File
		result = append(result, reMap)
	}
	reMap["data"] = result
	commonResponse.OkWithData(ctx, reMap)

}

// CreateMessage 新增留言
//
//	@Summary		新增留言
//	@Description	新增留言。
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			x-token		header		string			true	"认证令牌"
//	@Param			categoryId	path		int				true	"分类ID"
//	@Success		200			{object}	utils.Response	"地址列表获取成功"
//	@Failure		400			{object}	utils.Response	"无效的请求参数"
//	@Failure		404			{object}	utils.Response	"分类未找到"
//	@Failure		500			{object}	utils.Response	"服务器错误"
//	@Router			/v1/userop/message/create [post]
func CreateMessage(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	messageForm := forms.MessageForm{}
	if err := ctx.ShouldBindJSON(&messageForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.MessageSrvClient.CreateMessage(context.Background(), &commonpb.MessageRequest{
		MessageType: messageForm.MessageType,
		Message:     messageForm.Message,
		Subject:     messageForm.Subject,
		File:        messageForm.File,
		UserId:      int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("添加留言失败")
		HandleGrpcErrorToHttp(err, ctx, "用户操作srv")
		return
	}
	reMap := make(map[string]interface{})
	reMap["id"] = rsp.Id
	ctx.JSON(http.StatusOK, reMap)
}
