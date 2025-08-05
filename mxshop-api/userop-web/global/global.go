package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/config"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Trans            ut.Translator
	ServerConfig     = &config.ServerConfig{}
	Logger           *zap.Logger
	UserFavSrvClient proto.UserFavClient
	MessageSrvClient proto.MessageClient
	AddressSrvClient proto.AddressClient
	GoodsSrvClient   proto.GoodsClient
	UserOpConn       *grpc.ClientConn
	GoodsConn        *grpc.ClientConn
)
