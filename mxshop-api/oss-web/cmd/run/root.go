package run

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	commonInitialize "github.com/zhengpanone/mxshop/mxshop-api/common/initialize"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/common/utils/register/consul"
	"github.com/zhengpanone/mxshop/mxshop-api/oss-web/global"
	"github.com/zhengpanone/mxshop/mxshop-api/oss-web/initialize"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var CmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run oss-web server",
	Run:   runFunction,
}

var (
	configPath string
	crontab    string
	mode       string
)

func init() {
	CmdRun.Flags().StringVarP(&configPath, "config path", "c", "", "config path")
	CmdRun.Flags().StringVarP(&mode, "mode", "m", "debug", "debug or release")
	CmdRun.Flags().StringVarP(&crontab, "cron", "t", "open", "scheduled task control open or close")
}

func runFunction(cmd *cobra.Command, args []string) {
	//var err error
	// 判断是否编译线上版本
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}

	// 1.初始化配置文件
	initialize.InitConfig(configPath)

	// 2.初始化Logger
	logConfig := global.ServerConfig.LogConfig
	err := commonInitialize.InitLogger(logConfig.Filename, logConfig.MaxSize, logConfig.MaxBackups, logConfig.MaxAge, logConfig.Level)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger:%v", err))
	}
	global.Logger = commonInitialize.GetLogger()
	zap.ReplaceGlobals(global.Logger)
	global.Logger.Info("日志初始化成功")

	// 初始化oss
	initialize.InitOSS(global.ServerConfig.OssInfo)
	// 3.初始化routers
	Router := initialize.Routers()
	//Router.Use(gin.Logger())
	registerMiddleware(Router) //注册全局中间件

	// 4.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Errorf("初始化翻译器错误")
		return
	}

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
	// 服务注册
	registerClient := consul.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	serviceId := uuid.NewV4().String()
	err = registerClient.Register(commonUtils.GetIP(), global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("oss服务注册失败：", err.Error())
	}

	global.Logger.Info("oss服务oss-web服务注册到注册中心")

	global.Logger.Info(fmt.Sprintf("启动oss服务器，访问地址：http://%s:%d", commonUtils.GetIP(), global.ServerConfig.Port))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", global.ServerConfig.Port),
		Handler: Router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Panic(fmt.Sprintf("用户服务器启动失败：%s", err.Error()))
		}
	}()
	shutdown(server, serviceId, registerClient)

}

// registerMiddleware 注册中间件
func registerMiddleware(r *gin.Engine) {
	// 打印日志 、异常保护
	//r.Use(commonMiddleware.GinLogger(global.Logger), commonMiddleware.GinRecovery(global.Logger, true))
	r.Use(commonMiddleware.Cors() /*, middleware.Trace()*/)
}

// shutdown 听信号, 执行关机操作
// param srv 需要关闭的Http Server实例
func shutdown(server *http.Server, serviceId string, registerClient consul.RegistryClient) {
	// 接收终止信号
	quit := make(chan os.Signal, 1)
	// 注册要捕获的信号，这里包括 Ctrl+C（SIGINT）和终止信号（SIGTERM）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	global.Logger.Info("Waiting for Ctrl+C (SIGINT)...")
	// 阻塞等待信号
	sig := <-quit
	global.Logger.Info(fmt.Sprintf("Received signal: %v\n", sig))

	// 创建context 并设置超时时间, 确保服务关闭的操作在指定时间内完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 在注册中心注销服务
	if err := registerClient.DeRegister(serviceId); err != nil {
		global.Logger.Info(fmt.Sprintf("oss-web 在注册中心注销失败：%v\n", err.Error()))
	}
	global.Logger.Info("oss-web 在注册中心注销成功：")
	// 调用Http实例的Shutdown方法 关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Fatal("oss-web 关闭错误", zap.Error(err))
	}
	global.Logger.Info("oss-web 关闭")
}
