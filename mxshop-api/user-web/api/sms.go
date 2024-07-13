package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"user-web/forms"
	"user-web/utils"
)

func SendSms(ctx *gin.Context) {
	smsForm := forms.SendSmsForm{}
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
