package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/config"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Trans          ut.Translator
	ServerConfig   = &config.ServerConfig{}
	GoodsSrvClient proto.GoodsClient
	GoodsConn      *grpc.ClientConn
	Logger         *zap.Logger
)
