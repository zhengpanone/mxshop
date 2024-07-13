package global

import (
	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc"
	"order-web/config"
	"order-web/proto"
)

var (
	Trans              ut.Translator
	ServerConfig       *config.ServerConfig = &config.ServerConfig{}
	OrderSrvClient     proto.OrderClient
	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
	OrderConn          *grpc.ClientConn
	GoodsConn          *grpc.ClientConn
	InventoryConn      *grpc.ClientConn
)
