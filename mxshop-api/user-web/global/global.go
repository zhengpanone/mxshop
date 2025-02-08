package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/zhengpanone/mxshop/user-web/config"
	"github.com/zhengpanone/mxshop/user-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Trans         ut.Translator
	ServerConfig  = &config.ServerConfig{}
	Logger        *zap.Logger
	RedisClient   *redis.Client
	UserSrvClient proto.UserClient
	UserConn      *grpc.ClientConn
)
