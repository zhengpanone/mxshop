package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Trans         ut.Translator
	ServerConfig  = &config.ServerConfig{}
	Logger        *zap.Logger
	UserSrvClient commonpb.UserClient
	UserConn      *grpc.ClientConn
	RedisClient   *redis.Client
)
