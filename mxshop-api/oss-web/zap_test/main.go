package main

import (
	"go.uber.org/zap"
)

func main() {
	//logger, _ := zap.NewProduction() // 生成环境
	logger, _ := zap.NewDevelopment() //开发环境
	defer logger.Sync()
	url := "https://imooc.com"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL", "url", url, "attempt", 3)
	sugar.Infof("Failed to fetch URL:%s", url)

	logger.Info("failed to fetch url", zap.String("url", url), zap.Int("nums", 3))
}
