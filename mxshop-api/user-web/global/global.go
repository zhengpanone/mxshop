package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/config"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Trans         ut.Translator
	ServerConfig  = &config.ServerConfig{}
	Logger        *zap.Logger
	UserSrvClient proto.UserClient
	UserConn      *grpc.ClientConn
)
