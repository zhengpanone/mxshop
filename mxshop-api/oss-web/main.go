package main

import (
	commonInitialize "common/initialize"
	commonMiddleware "common/middleware"
	commonUtils "common/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"oss-web/global"
	"oss-web/initialize"
	"oss-web/utils"
	"oss-web/utils/register/consul"

	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {

	// 1.初始化配置文件
	initialize.InitConfig()

	// 2.初始化Logger
	logConfig := global.ServerConfig.LogConfig
	zapLogger, err := commonInitialize.InitLogger(logConfig.Filename, logConfig.MaxSize, logConfig.MaxBackups, logConfig.MaxAge, logConfig.Level)
	if err != nil {
		panic(err)
	}
	global.Logger = zapLogger
	zap.ReplaceGlobals(zapLogger)
	zap.L().Info("日志初始化成功")

	// 初始化oss
	initialize.InitOSS(global.ServerConfig.OssInfo)
	// 3.初始化routers
	Router := initialize.Routers()
	Router.Use(gin.Logger())
	// 4.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Errorf("初始化翻译器错误")
		return
	}

	registerMiddleware(Router) //注册全局中间件

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("run exe dir is %v\n", dir)
	//
	currentMod := gin.Mode()
	if currentMod == gin.ReleaseMode {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}
	// 服务注册
	registerClient := consul.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	serviceId := uuid.NewV4().String()
	err = registerClient.Register(commonUtils.GetIP(), global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("oss服务注册失败：", err.Error())
	}
	global.Logger.Info(fmt.Sprintf("oss服务oss-web服务注册到注册中心"))

	global.Logger.Info(fmt.Sprintf("启动oss服务器，访问地址：http://%s:%d", commonUtils.GetIP(), global.ServerConfig.Port))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.ServerConfig.Port),
		Handler: Router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Panic(fmt.Sprintf("用户服务器启动失败：%s", err.Error()))
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal, 1)
	// 注册要捕获的信号，这里包括 Ctrl+C（SIGINT）和终止信号（SIGTERM）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	global.Logger.Info("Waiting for Ctrl+C (SIGINT)...")
	// 阻塞等待信号
	sig := <-quit
	global.Logger.Info(fmt.Sprintf("Received signal: %v", sig))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Fatal("Server Shutdown:", zap.Error(err))
	}
	if err = registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败：", err.Error())
	}
	global.Logger.Info("注销成功：")
}

// registerMiddleware 注册中间件
func registerMiddleware(r *gin.Engine) {
	// 打印日志 、异常保护
	r.Use(commonMiddleware.GinLogger(global.Logger), commonMiddleware.GinRecovery(global.Logger, true))
}
