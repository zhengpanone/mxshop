package global

import (
	ut "github.com/go-playground/universal-translator"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Trans              ut.Translator
	ServerConfig       *config.ServerConfig = &config.ServerConfig{}
	Logger             *zap.Logger
	OrderSrvClient     commonpb.OrderClient
	GoodsSrvClient     commonpb.GoodsClient
	InventorySrvClient commonpb.InventoryClient
	OrderConn          *grpc.ClientConn
	GoodsConn          *grpc.ClientConn
	InventoryConn      *grpc.ClientConn
)
