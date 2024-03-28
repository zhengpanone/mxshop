package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
	"mxshop-api/user-web/middlewares"
	myValidator "mxshop-api/user-web/validator"
	"os"
	"path/filepath"
)

func main() {
	// 1.初始化Logger
	initialize.InitLogger()
	// 2.初始化配置文件
	initialize.InitConfig()
	// 3. 初始化连接redis
	if global.ServerConfig.System.UseRedis {
		initialize.InitRedis()
	}
	// 3.初始化routers
	Router := initialize.Routers()
	// 4.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Errorf("初始化翻译器错误")
		return
	}

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myValidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	Router.Use(middlewares.MyLogger()) //注册全局中间件

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("run exe dir is %v", dir)
	port := global.ServerConfig.Port
	zap.S().Debugf("启动服务器，访问地址：http://127.0.0.1:%d", port)
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("服务器启动失败：", err.Error())
	}

}
