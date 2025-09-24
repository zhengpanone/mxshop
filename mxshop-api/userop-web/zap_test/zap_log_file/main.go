package main

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./myprojectlog", // 输出到文件
		"stderr",         // 标准错误输出
		"stdout",
	}
	cfg.ErrorOutputPaths = []string{
		"./myprojectlog.errs", // 输出到文件
		"stderr",              // 标准输出
	}
	return cfg.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
		//panic("初始化logger失败")
	}
	sugar := logger.Sugar()
	defer sugar.Sync()
	url := "https://imooc.com"
	sugar.Infow("failed to fetch URL", "url", url, "attempt", 3)
	sugar.Error("failed to fetch URL", "url", url, "attempt", 3)
	sugar.Infof("Failed to fetch URL:%s", url)

	logger.Info("failed to fetch url", zap.String("url", url), zap.Int("nums", 3))
}
