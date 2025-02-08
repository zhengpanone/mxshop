package controller

import (
	"github.com/gin-gonic/gin"
	commonUtils "github.com/zhengpanone/mxshop/common/utils"
	"github.com/zhengpanone/mxshop/user-web/forms"
	"github.com/zhengpanone/mxshop/user-web/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func GenerateCaptcha(ctx *gin.Context) {
	captchaForm := forms.CaptchaForm{}
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
		"captchaId": id,
		"picPath":   base64,
	}
	commonUtils.OkWithData(ctx, h)
}

func HandleGrpcErrorToHttp(err error, c *gin.Context, srvName string) {
	// 将grpc的code转换成http的code
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				commonUtils.ErrorWithCodeAndMsg(c, http.StatusNotFound, srvName+":"+e.Message())
			case codes.Internal:
				commonUtils.ErrorWithCodeAndMsg(c, http.StatusInternalServerError, srvName+":"+"内部错误"+e.Message())
			case codes.InvalidArgument:
				commonUtils.ErrorWithCodeAndMsg(c, http.StatusBadRequest, srvName+":"+"参数错误"+e.Message())
			case codes.Unavailable:
				commonUtils.ErrorWithCodeAndMsg(c, http.StatusInternalServerError, srvName+":"+"不可用")
			default:
				commonUtils.ErrorWithCodeAndMsg(c, http.StatusInternalServerError, srvName+":"+"其他错误"+e.Message())
			}

		}

	}
}
