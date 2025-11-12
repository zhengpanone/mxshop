package initialize

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/global"
	"go.uber.org/zap"

	"strings"
)

var RootPath string

func GetConfigFromNacos() {
	nacosConfig := global.ServerConfig.Nacos

	// 详细输出配置信息用于调试
	global.Logger.Info("=== Nacos配置参数详情 ===",
		zap.String("Host", nacosConfig.Host),
		zap.Uint64("Port", nacosConfig.Port),
		zap.String("Namespace", nacosConfig.Namespace),
		zap.String("DataId", nacosConfig.DataId),
		zap.String("Group", nacosConfig.Group),
		zap.String("ContextPath", nacosConfig.ContextPath),
		zap.String("User", nacosConfig.User),
		zap.String("Password", nacosConfig.Password))
	// 2. 创建 Nacos 服务端配置
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			nacosConfig.Host,
			nacosConfig.Port,
			constant.WithContextPath(nacosConfig.ContextPath),
			constant.WithScheme("http"),
		),
	}
	tempDir := os.TempDir()
	nacosLogDir := filepath.Join(tempDir, "nacos", "log")
	nacosCacheDir := filepath.Join(tempDir, "nacos", "cache")

	// 确保目录存在
	if err := os.MkdirAll(nacosLogDir, 0755); err != nil {
		global.Logger.Warn("创建日志目录失败", zap.Error(err))
	}
	if err := os.MkdirAll(nacosCacheDir, 0755); err != nil {
		global.Logger.Warn("创建缓存目录失败", zap.Error(err))
	}

	global.Logger.Info("Nacos目录配置",
		zap.String("logDir", nacosLogDir),
		zap.String("cacheDir", nacosCacheDir))

	// 4. 创建 Nacos 客户端配置
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(nacosConfig.Namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(nacosLogDir),
		constant.WithCacheDir(nacosCacheDir),
		constant.WithLogLevel("debug"),
		constant.WithUsername(nacosConfig.User),
		constant.WithPassword(nacosConfig.Password),
	)

	// 5. 创建 Nacos 客户端
	global.Logger.Info("正在创建Nacos客户端...")
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		global.Logger.Error("创建Nacos客户端失败", zap.Error(err))
		panic(fmt.Sprintf("创建Nacos客户端失败: %v", err))
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
	})
	if err != nil {
		global.Logger.Error("获取Nacos配置失败",
			zap.String("dataId", nacosConfig.DataId),
			zap.String("group", nacosConfig.Group),
			zap.String("namespace", nacosConfig.Namespace),
			zap.Error(err))
		panic(fmt.Sprintf("获取Nacos配置失败: %v", err))
	}
	global.Logger.Info("✓ 配置获取成功")
	// 将配置内容设置到 Viper
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(content)); err != nil {
		log.Fatalf("nacos配置设置到viper中失败：%v", err)
	}
	if err := viper.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}

	/*config := global.ServerConfig
	errs = yaml.Unmarshal([]byte(content), &config)
	if errs != nil {
		zap.S().Fatalf("读取nacos配置文件，转换yaml失败:%s", errs)
	}
	fmt.Printf("Decoded YAML: %+v\\n", config)*/

	// 监听 Nacos 配置变化。
	err = client.ListenConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			global.Logger.Info("配置文件变化-------")
			//fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)

			/*errs = yaml.Unmarshal([]byte(content), &global.ServerConfig)
			if errs != nil {
				zap.S().Fatalf("读取nacos配置文件，转换yaml失败:%s", errs)
			}*/
			// 当配置变化时，更新 Viper 中的配置。
			if err := viper.ReadConfig(strings.NewReader(data)); err != nil {
				global.Logger.Info("Viper read config error: %v\n", zap.Error(err))
			} else {
				global.Logger.Info("Config has been updated.")
				if err := viper.Unmarshal(global.ServerConfig); err != nil {
					panic(err)
				}
				//fmt.Printf("----2---%v\n", global.ServerConfig)
				InitSrvConn()
			}
		},
	})
}

func InitConfig(path string) {
	if path == "" {
		RootPath = "."
	} else {
		RootPath = path
	}
	mode := gin.Mode()
	configFilePrefix := "config"

	configFileName := fmt.Sprintf("%s-pro.yaml", configFilePrefix)
	if mode == gin.DebugMode {
		configFileName = fmt.Sprintf("%s-dev.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(filepath.Join(RootPath, configFileName))

	if err := v.ReadInConfig(); err != nil {
		global.Logger.Info(fmt.Sprintf("Error reading config file, %s", err))
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	global.Logger.Info(fmt.Sprintf("配置信息：%v", global.ServerConfig))
	// 从nacos中读取配置信息
	GetConfigFromNacos()
	global.Logger.Info(fmt.Sprintf("配置信息：%v", global.ServerConfig))
	// viper的功能-动态监控变化

	// 监听本地配置文件的变化（可选）
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		global.Logger.Info(fmt.Sprintf("Local Config File Changed:%e", e.Name))
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		InitSrvConn()
	})

	/*config, errs := configClient.GetConfig(vo.ConfigParam{
		DataId: "your-data-id",
		Group:  "your-group",
	})
	if errs != nil {
		log.Fatal(errs)
	}*/
	/*viper.SetConfigType("yaml")
	bytes.NewBuffer()
	if errs := viper.ReadConfig(bytes.NewBuffer(config)); errs != nil {
		log.Fatal(errs)
	}
	fmt.Println("Config updated:", viper.AllSettings())*/

}
