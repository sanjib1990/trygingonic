package Logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"trygonic/app/config"
)

var logger *zap.Logger

func Get() *zap.Logger {
	return logger
}

func init() {
	// First, define our level-handling logic.
	_level := config.Values.LogLevel
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level(_level)
	})

	core := zapcore.NewTee(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		priority))

	logger = zap.New(core)
	logger.Info("Logger initiated... LEVEL: " + config.Values.LogLevel)
}

func level(level string) zapcore.Level {
	switch level {
	case "INFO":
		return zapcore.InfoLevel
	case "DEBUG":
		return zapcore.DebugLevel
	case "WARN":
		return zapcore.WarnLevel
	default:
		return zapcore.ErrorLevel
	}
}
