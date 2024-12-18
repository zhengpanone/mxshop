package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"user-web/forms"
	"user-web/utils"
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
	utils.OkWithData(ctx, h)
}

func HandleGrpcErrorToHttp(err error, c *gin.Context, srvName string) {
	// 将grpc的code转换成http的code
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				utils.ErrorWithCodeMsg(c, http.StatusNotFound, srvName+":"+e.Message())
			case codes.Internal:
				utils.ErrorWithCodeMsg(c, http.StatusInternalServerError, srvName+":"+"内部错误"+e.Message())
			case codes.InvalidArgument:
				utils.ErrorWithCodeMsg(c, http.StatusBadRequest, srvName+":"+"参数错误"+e.Message())
			case codes.Unavailable:
				utils.ErrorWithCodeMsg(c, http.StatusInternalServerError, srvName+":"+"不可用")
			default:
				utils.ErrorWithCodeMsg(c, http.StatusInternalServerError, srvName+":"+"其他错误"+e.Message())
			}

		}

	}
}
