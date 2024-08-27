package global

import (
	commonConfig "common/config"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"oss-web/config"
)

var (
	Trans        ut.Translator
	ServerConfig = &config.ServerConfig{}
	Logger       *zap.Logger
	NacosConfig  = &commonConfig.NacosConfig{}
)
