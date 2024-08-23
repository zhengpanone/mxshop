package global

import (
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	commonConfig "mxshop-api/common/config"
	"oss-web/config"
)

var (
	Trans        ut.Translator
	ServerConfig = &config.ServerConfig{}
	Logger       *zap.Logger
	NacosConfig  = &commonConfig.NacosConfig{}
)
