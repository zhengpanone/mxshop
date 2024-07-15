package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"user-web/config"
	"user-web/proto"
)

var (
	Trans         ut.Translator
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	RedisClient   *redis.Client
	UserSrvClient proto.UserClient
	UserConn      *grpc.ClientConn
)
