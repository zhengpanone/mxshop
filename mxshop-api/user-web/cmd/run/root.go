package run

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	commonInitialize "github.com/zhengpanone/mxshop/mxshop-api/common/initialize"
	commonMiddleware "github.com/zhengpanone/mxshop/mxshop-api/common/middleware"
	commonUtils "github.com/zhengpanone/mxshop/mxshop-api/common/utils"
	"github.com/zhengpanone/mxshop/mxshop-api/common/utils/register/consul"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/global"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/initialize"
	myValidator "github.com/zhengpanone/mxshop/mxshop-api/user-web/validator"
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
	Short: "Run user-web server",
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
	//var errs error
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
	zap.L().Info("日志初始化成功")

	// 3. 初始化连接redis
	if global.ServerConfig.System.UseRedis {
		commonInitialize.InitRedis(global.ServerConfig.RedisConfig)
	}
	// 4.初始化routers
	Router := initialize.Routers()
	// 5.初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Errorf("初始化翻译器错误")
		return
	}
	// 6. 初始化srv连接
	initialize.InitSrvConn()

	// 7. 注册验证器
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
			global.ServerConfig.Port = uint32(port)
		}
	}

	registerClient := consul.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	serviceId := uuid.NewV4().String()
	err = registerClient.Register(commonUtils.GetIP(), global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)

	if err != nil {
		zap.S().Panic("用户服务注册失败：", err.Error())
	}
	global.Logger.Info("用户服务user-web服务注册到注册中心")
	zap.S().Debugf("启动用户服务器，访问地址：http://%s:%d", commonUtils.GetIP(), global.ServerConfig.Port)
	zap.S().Info(fmt.Sprintf("swagger，访问地址：http://%s:%d/swagger/index.html", commonUtils.GetIP(), global.ServerConfig.Port))
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.ServerConfig.Port),
		Handler: Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Panic("用户服务器启动失败：", err.Error())
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
	global.Logger.Info("Waiting for Ctrl+C (SIGINT)...")
	// 阻塞等待信号
	sig := <-quit
	global.Logger.Info(fmt.Sprintf("Received signal: %v\n", sig))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 在注册中心注销服务
	if err := registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("用户服务goods-web 在注册中心注销失败：", err.Error())
	}
	zap.S().Info("用户服务goods-web 在注册中心注销成功")

	// 调用Http实例的Shutdown方法 关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Fatal("用户服务goods-web 关闭错误", zap.Error(err))
	}

	global.Logger.Info("商品服务goods-web 关闭")
}
