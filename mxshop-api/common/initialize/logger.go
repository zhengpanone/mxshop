package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// InitLogger 初始化日志
// 如果需要根据环境打印不同的日志, 也可以根据配置文件中的mode 来作为判断条件, 在这个函数的入参里面增加一个mode, 将配置文件里的数据传进来
func InitLogger(filename string, maxSize, maxBackup, maxAge int, level string) error {
	var err error
	writerSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err = l.UnmarshalText([]byte(level)); err != nil {
		return err
	}
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writerSyncer, l),
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), l),
	)
	// 初始化一个全局对象, 并添加调用栈信息
	once.Do(func() {
		logger = zap.New(core, zap.AddCaller())
		if err != nil {
			logger = nil
		}

	})
	return err
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetLogger() *zap.Logger {
	if logger == nil {
		panic("Logger not initialized. Call InitLogger first!")
	}
	return logger
}
