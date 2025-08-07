package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"sync"
)

var (
	Logger      *zap.Logger
	Once        sync.Once
	RedisClient *redis.Client
)
