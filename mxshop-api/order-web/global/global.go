package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/zhengpanone/mxshop/order-web/config"
	"github.com/zhengpanone/mxshop/order-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Trans              ut.Translator
	ServerConfig       *config.ServerConfig = &config.ServerConfig{}
	Logger             *zap.Logger
	OrderSrvClient     proto.OrderClient
	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
	OrderConn          *grpc.ClientConn
	GoodsConn          *grpc.ClientConn
	InventoryConn      *grpc.ClientConn
)
