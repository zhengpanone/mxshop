package main

import (
	commonInitialize "common/initialize"
	commonMiddleware "common/middleware"
	commonUtils "common/utils"
	"common/utils/register/consul"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"order-web/global"
	"order-web/initialize"

	myValidator "order-web/validator"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	// 1.初始化配置文件
	initialize.InitConfig()
	// 2.初始化Logger
	logConfig := global.ServerConfig.LogConfig
	err := commonInitialize.InitLogger(logConfig.Filename, logConfig.MaxSize, logConfig.MaxBackups, logConfig.MaxAge, logConfig.Level)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger:%v", err))
	}
	global.Logger = commonInitialize.GetLogger()
	zap.ReplaceGlobals(global.Logger)
	global.Logger.Info("日志初始化成功")

	// 3.初始化routers
	Router := initialize.Routers()
	// 4.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Errorf("初始化翻译器错误")
		return
	}
	// 5. 初始化srv连接
	initialize.InitSrvConn()

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

	registerMiddleware(Router) //注册全局中间件

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("run exe dir is %v\n", dir)
	//
	currentMod := gin.Mode()
	if currentMod == gin.ReleaseMode {
		port, err := commonUtils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}
	registerClient := consul.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	serviceId := uuid.NewV4().String()
	err = registerClient.Register(commonUtils.GetIP(), global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("订单服务注册失败：", err.Error())
	}
	global.Logger.Info(fmt.Sprintf("订单服务order-web服务注册到注册中心"))
	global.Logger.Info(fmt.Sprintf("启动订单服务器，访问地址：http://%s:%d", commonUtils.GetIP(), global.ServerConfig.Port))
	global.Logger.Info(fmt.Sprintf("swagger，访问地址：http://%s:%d/swagger/index.html", commonUtils.GetIP(), global.ServerConfig.Port))
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("订单服务启动失败：", err.Error())
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal, 1)
	// 注册要捕获的信号，这里包括 Ctrl+C（SIGINT）和终止信号（SIGTERM）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Waiting for Ctrl+C (SIGINT)...")
	// 阻塞等待信号
	sig := <-quit
	fmt.Printf("Received signal: %v\n", sig)
	if err = registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败：", err.Error())
	}
	zap.S().Info("注销成功：")
}

// registerMiddleware 注册中间件
func registerMiddleware(r *gin.Engine) {
	// 打印日志 、异常保护
	r.Use(commonMiddleware.GinLogger(global.Logger), commonMiddleware.GinRecovery(global.Logger, true))

}
