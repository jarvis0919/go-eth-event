package initialize

import (
	"event/global"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog() {
	flushLog()
	go func() {
		for {
			flushLog()
			// 等到下一个小时
			nextHour := time.Now().Truncate(time.Hour).Add(time.Hour)
			<-time.After(time.Until(nextHour))
		}
	}()
}

func flushLog() {
	// 获取当前小时并生成新日志文件名
	now := time.Now()
	logFilename := fmt.Sprintf("./log/app_%s.log", now.Format("2006-01-02-15"))
	// 配置lumberjack Logger
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFilename, // 使用基于日期的文件名
		MaxSize:    10,          // 每个日志文件的最大 MB 数
		MaxBackups: 3,           // 保留旧文件的最大个数
		MaxAge:     28,          // 保留旧文件的最大天数
		Compress:   true,        // 是否压缩
	}
	// 控制台输出
	consoleEncoder := zapcore.NewConsoleEncoder(getConfig(zapcore.CapitalColorLevelEncoder))
	consoleDebugging := zapcore.Lock(os.Stdout) // 输出到控制台
	// 日志文件输出
	fileEncoder := zapcore.NewJSONEncoder(getConfig(zapcore.CapitalLevelEncoder))
	fileSyncer := zapcore.AddSync(lumberjackLogger)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebugging, zap.InfoLevel),
		zapcore.NewCore(fileEncoder, fileSyncer, zap.InfoLevel),
	)
	// 创建新的日志记录器
	global.LOG = zap.New(core, zap.AddCaller())
}

func getConfig(color zapcore.LevelEncoder) zapcore.EncoderConfig {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = color
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		type appendTimeEncoder interface {
			AppendTimeLayout(time.Time, string)
		}
		if enc, ok := enc.(appendTimeEncoder); ok {
			enc.AppendTimeLayout(t, "[2006-01-02 15:04:05]")
			return
		}

		enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
	}
	return encoderConfig
}
