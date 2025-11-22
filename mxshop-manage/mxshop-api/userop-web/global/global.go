package global

import (
	ut "github.com/go-playground/universal-translator"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Trans            ut.Translator
	ServerConfig     = &config.ServerConfig{}
	Logger           *zap.Logger
	UserFavSrvClient commonpb.UserFavClient
	MessageSrvClient commonpb.MessageClient
	AddressSrvClient commonpb.AddressClient
	GoodsSrvClient   commonpb.GoodsClient
	UserOpConn       *grpc.ClientConn
	GoodsConn        *grpc.ClientConn
)
