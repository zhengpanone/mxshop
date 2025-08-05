package global

import (
	ut "github.com/go-playground/universal-translator"
	commonConfig "github.com/zhengpanone/mxshop/mxshop-api/common/config"
	"github.com/zhengpanone/mxshop/mxshop-api/oss-web/config"
	"go.uber.org/zap"
)

var (
	Trans        ut.Translator
	ServerConfig = &config.ServerConfig{}
	Logger       *zap.Logger
	NacosConfig  = &commonConfig.NacosConfig{}
	OSSClient    = &config.OSSClients{}
	OSSConfig    = &config.OssConfig{}
)
