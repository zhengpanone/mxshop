package global

import (
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"userop-web/config"
	"userop-web/proto"
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
