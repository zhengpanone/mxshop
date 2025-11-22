package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/zhengpanone/mxshop/mxshop-api/common/claims"
	"go.uber.org/zap"
	"sync"
)

var (
	Logger       *zap.Logger
	Once         sync.Once
	RedisClient  *redis.Client
	TokenManager *claims.TokenManager
)
