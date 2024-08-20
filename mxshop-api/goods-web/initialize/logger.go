package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goods-web/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// InitLogger 初始化日志
// 如果需要根据环境打印不同的日志, 也可以根据配置文件中的mode 来作为判断条件, 在这个函数的入参里面增加一个mode, 将配置文件里的数据传进来
func InitLogger(filename string, maxSize, maxBackup, maxAge int, level string) (err error) {
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
	global.Logger = zap.New(core, zap.AddCaller())
	// 替换zap包全局的logger
	zap.ReplaceGlobals(global.Logger)
	zap.L().Info("日志初始化成功")
	return
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
