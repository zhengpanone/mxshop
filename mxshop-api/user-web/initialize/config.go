package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"strings"
	"user-web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func GetConfigFromNacos() {
	serverConfig := global.ServerConfig
	//创建 Nacos 服务端配置
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(serverConfig.Nacos.Host, serverConfig.Nacos.Port, constant.WithContextPath(serverConfig.Nacos.ContextPath)),
	}
	//创建 Nacos 客户端配置
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(serverConfig.Nacos.Namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 创建 Nacos 客户端
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: serverConfig.Nacos.DataId,
		Group:  serverConfig.Nacos.Group,
	})
	// 将配置内容设置到 Viper
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(content)); err != nil {
		log.Fatalf("nacos配置设置到viper中失败：%v", err)
	}
	if err := viper.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	fmt.Printf("----1---%v\n", global.ServerConfig)

	/*config := global.ServerConfig
	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		zap.S().Fatalf("读取nacos配置文件，转换yaml失败:%s", err)
	}
	fmt.Printf("Decoded YAML: %+v\\n", config)*/

	// 监听 Nacos 配置变化。
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "user-web.yaml",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件变化-------")
			//fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)

			/*err = yaml.Unmarshal([]byte(content), &global.ServerConfig)
			if err != nil {
				zap.S().Fatalf("读取nacos配置文件，转换yaml失败:%s", err)
			}*/
			// 当配置变化时，更新 Viper 中的配置。
			if err := viper.ReadConfig(strings.NewReader(data)); err != nil {
				fmt.Printf("Viper read config error: %v\n", err)
			} else {
				fmt.Println("Config has been updated.")
				if err := viper.Unmarshal(global.ServerConfig); err != nil {
					panic(err)
				}
				fmt.Printf("----2---%v\n", global.ServerConfig)
				InitSrvConn()
			}
		},
	})
}

func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"

	configFileName := fmt.Sprintf("./%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("./%s-dev.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFileName)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	//zap.S().Infof("配置信息：%v", global.ServerConfig)
	// 从nacos中读取配置信息
	GetConfigFromNacos()
	zap.S().Infof("配置信息：%v", global.ServerConfig)
	// viper的功能-动态监控变化

	// 监听本地配置文件的变化（可选）
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		zap.S().Infof("Local Config File Changed:%e", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		InitSrvConn()
	})

	/*config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "your-data-id",
		Group:  "your-group",
	})
	if err != nil {
		log.Fatal(err)
	}*/
	/*viper.SetConfigType("yaml")
	bytes.NewBuffer()
	if err := viper.ReadConfig(bytes.NewBuffer(config)); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config updated:", viper.AllSettings())*/

}
