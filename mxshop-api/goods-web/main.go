package main

import (
	commonInitialize "common/initialize"
	commonMiddleware "common/middleware"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"goods-web/global"
	"goods-web/initialize"
	"goods-web/utils"
	"goods-web/utils/register/consul"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

// @title 商品服务
// @description 慕学商城项目
// @version 1.0
// @contact.name zhengpanone
// @contact.url http://.....
// @host localhost:8080
// @BasePath

func main() {
	// 1.初始化配置文件
	initialize.InitConfig()
	// 2.初始化Logger
	logConfig := global.ServerConfig.LogConfig
	logger, err := commonInitialize.InitLogger(logConfig.Filename, logConfig.MaxSize, logConfig.MaxBackups, logConfig.MaxAge, logConfig.Level)
	if err != nil {
		panic(err)
	}
	global.Logger = logger
	zap.ReplaceGlobals(logger)
	zap.L().Info("日志初始化成功")

	// 3.初始化routers
	Router := initialize.Routers()
	// 4.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Errorf("初始化翻译器错误")
		return
	}
	// 5. 初始化srv连接
	initialize.InitSrvConn()

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
	// 注册到注册中心
	registerClient := consul.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	serviceId := uuid.NewV4().String()
	err = registerClient.Register(utils.GetIP(), global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)

	if err != nil {
		zap.S().Panic("商品服务goods-web 注册失败：", err.Error())
	}

	global.Logger.Info(fmt.Sprintf("商品服务goods-web服务注册到注册中心"))

	global.Logger.Info(fmt.Sprintf("启动商品服务goods-web，访问地址：http://%s:%d", utils.GetIP(), global.ServerConfig.Port))
	global.Logger.Info(fmt.Sprintf("swagger，访问地址：http://%s:%d/swagger/index.html", utils.GetIP(), global.ServerConfig.Port))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", global.ServerConfig.Port),
		Handler: Router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Panic("商品服务goods-web 启动失败：", err)
		}
	}()

	shutdown(server, serviceId, registerClient)
}

// registerMiddleware 注册中间件
func registerMiddleware(r *gin.Engine) {
	// 打印日志 、异常保护
	r.Use(commonMiddleware.GinLogger(global.Logger), commonMiddleware.GinRecovery(global.Logger, true))

}

// shutdown 听信号, 执行关机操作
// param srv 需要关闭的Http Server实例
func shutdown(server *http.Server, serviceId string, registerClient consul.RegistryClient) {
	// 接收终止信号
	quit := make(chan os.Signal, 1)
	// 注册要捕获的信号，这里包括 Ctrl+C（SIGINT）和终止信号（SIGTERM）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Waiting for Ctrl+C (SIGINT)...")

	// 阻塞等待信号 启动之后 这里会阻塞, 等到退出的时候(通道里面有值)才会往下执行
	sig := <-quit
	global.Logger.Info(fmt.Sprintf("Received signal: %v\n", sig))

	// 创建context 并设置超时时间, 确保服务关闭的操作在指定时间内完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 在注册中心注销服务
	if err := registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("商品服务goods-web 在注册中心注销失败：", err.Error())
	}
	zap.S().Info("商品服务goods-web 在注册中心注销成功：")

	// 调用Http实例的Shutdown方法 关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Fatal("商品服务goods-web 关闭错误", zap.Error(err))
	}
	global.Logger.Info("商品服务goods-web 关闭")
}
