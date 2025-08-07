package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	commonGlobal "github.com/zhengpanone/mxshop/mxshop-api/common/global"
	"github.com/zhengpanone/mxshop/mxshop-api/goods-web/global"
)

func InitRedis() {
	redisConfig := global.ServerConfig.RedisConfig
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		DB:       redisConfig.Database,
		Password: redisConfig.Password,
	})
	ctx := context.Background()
	//timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancelFunc()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic("redis初始化失败" + err.Error())
	}
	commonGlobal.RedisClient = redisClient
}
