package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/request"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/utils"
	"go.uber.org/zap"
	"net/http"
)

func SendSms(ctx *gin.Context) {
	smsForm := request.SendSmsForm{}
	if err := ctx.ShouldBind(&smsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	code := utils.GenerateSmsCode(5)
	if err := utils.SendSms(smsForm.Mobile, code); err != nil {
		zap.S().Errorf("发送短信失败:%s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "发送短信失败" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
