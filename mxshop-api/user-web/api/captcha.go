package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GenerateCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, base64, _, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误:,", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   base64,
	})
}

func VerifyCaptcha(id, answer string) bool {
	return store.Verify(id, answer, true)
}
