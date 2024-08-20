package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"user-web/config"
	"user-web/proto"
)

var (
	Trans         ut.Translator
	ServerConfig  = &config.ServerConfig{}
	Logger        *zap.Logger
	RedisClient   *redis.Client
	UserSrvClient proto.UserClient
	UserConn      *grpc.ClientConn
)
