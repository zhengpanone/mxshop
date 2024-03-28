package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"mxshop-api/user-web/global"
)

func InitRedis() {
	redisConfig := global.ServerConfig.RedisConfig
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		DB:       0,
		Password: "",
	})
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), global.ServerConfig.RedisConfig.DialTimeout)
	defer cancelFunc()
	_, err := redisClient.Ping(timeoutCtx).Result()
	if err != nil {
		panic("redis初始化失败" + err.Error())
	}
	global.GlobalRedisClient = redisClient
}
