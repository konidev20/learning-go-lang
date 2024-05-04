package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Create a lumberjack logger for log rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}

	// Create a syslog-compatible encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create a syslog-compatible encoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Create a zap core with the lumberjack logger and syslog encoder
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(lumberjackLogger),
		zap.InfoLevel,
	)

	// Create a new zap logger with the custom core
	logger := zap.New(core)
	defer logger.Sync()

	// Log messages using the logger
	logger.Info("This is an info message", zap.String("key", "value"))
	logger.Error("This is an error message", zap.Int("code", 500))
}
