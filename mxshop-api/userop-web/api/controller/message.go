package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	customClaims "github.com/zhengpanone/mxshop/common/claims"
	"github.com/zhengpanone/mxshop/userop-web/forms"
	"github.com/zhengpanone/mxshop/userop-web/global"
	"github.com/zhengpanone/mxshop/userop-web/proto"
	"go.uber.org/zap"
	"net/http"
)

func GetMessageList(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	claims, _ := ctx.Get("claims")
	// 管理员查询所有订单
	model := claims.(*customClaims.CustomClaims)
	request := proto.MessageRequest{}
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
	ctx.JSON(http.StatusOK, reMap)

}

func CreateMessage(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	messageForm := forms.MessageForm{}
	if err := ctx.ShouldBindJSON(&messageForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.MessageSrvClient.CreateMessage(context.Background(), &proto.MessageRequest{
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
