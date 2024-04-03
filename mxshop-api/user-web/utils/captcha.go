package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 创建store,保存验证码的位置,默认为mem(内存中)单机部署,如果要布置多台服务器,则可以设置保存在redis中
var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 获取验证码
func GenerateCaptcha(height, width int, length int) (string, string, error) {
	// 定义一个driver
	var driver base64Captcha.Driver
	// 创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driverString := base64Captcha.DriverString{
		Height:          height,                                 // 高度
		Width:           width,                                  // 宽度
		NoiseCount:      0,                                      // 干扰数
		ShowLineOptions: 2 | 4,                                  // 展示个数
		Length:          length,                                 // 长度
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm", //验证码随机字符串来源
		BgColor: &color.RGBA{ // 背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"}, // 字体
	}
	driver = driverString.ConvertFonts()
	//driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)

	// 生成验证码
	cp := base64Captcha.NewCaptcha(driver, store)
	id, base64, _, err := cp.Generate()

	return id, base64, err
	/*if err != nil {
		zap.S().Errorf("生成验证码错误:,", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   base64,
	})*/
}

func VerifyCaptcha(id, answer string) bool {
	return store.Verify(id, answer, true)
}
