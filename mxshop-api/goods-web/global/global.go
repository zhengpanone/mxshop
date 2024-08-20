package global

import (
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"goods-web/config"
	"goods-web/proto"
	"google.golang.org/grpc"
)

var (
	Trans          ut.Translator
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	GoodsSrvClient proto.GoodsClient
	GoodsConn      *grpc.ClientConn
	Logger         *zap.Logger
)
