// pkg/logger/logger.go
package logger

import (
	"context"

	"quanfuxia/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLogger *zap.Logger

// Init 初始化 logger
func Init() {
	c := config.Cfg.Log

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   c.File,
		MaxSize:    100, // MB
		MaxAge:     7,   // days
		MaxBackups: 30,
		Compress:   true,
	})

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	level := zapcore.DebugLevel
	if err := level.UnmarshalText([]byte(c.Level)); err != nil {
		level = zapcore.InfoLevel
	}

	core := zapcore.NewCore(encoder, writer, level)

	zapLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(zapLogger)
}

// L 返回全局 logger
func L() *zap.Logger {
	return zapLogger
}

// WithTrace 从 context 中取 trace_id 加入日志
func WithTrace(ctx context.Context) *zap.Logger {
	traceID, _ := ctx.Value("trace_id").(string)
	return zapLogger.With(zap.String("trace_id", traceID))
}
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, "trace_id", traceID)
}
