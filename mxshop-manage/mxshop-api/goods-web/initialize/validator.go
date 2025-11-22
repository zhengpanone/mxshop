package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/global"
	"reflect"
	"strings"
)

func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎，实现机制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "_" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() //英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语音环境
		uni := ut.New(enT, zhT, enT)

		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			entranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zhtranslations.RegisterDefaultTranslations(v, global.Trans)

		default:
			entranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return

}
