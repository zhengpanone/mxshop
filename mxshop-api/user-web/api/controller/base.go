package controller

import (
	"github.com/gin-gonic/gin"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/request"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/utils"
	"go.uber.org/zap"
	"net/http"
)

func GenerateCaptcha(ctx *gin.Context) {
	captchaForm := request.CaptchaForm{}
	if err := ctx.ShouldBind(&captchaForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	if captchaForm.Height == 0 {
		captchaForm.Height = 80
	}
	if captchaForm.Width == 0 {
		captchaForm.Width = 240
	}
	if captchaForm.Length == 0 {
		captchaForm.Length = 5
	}
	id, base64, err := utils.GenerateCaptcha(captchaForm.Height, captchaForm.Width, captchaForm.Length)
	if err != nil {
		zap.S().Errorf("生成验证码错误:,", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	h := gin.H{
		"captchaId":   id,
		"imageBase64": base64,
	}
	commonResponse.OkWithData(ctx, h)
}
