package logger

import (
	"context"
	"os"
	"strings"

	"github.com/arymaulanamalik/ai-engineer-playground/ai-playground/configs"
	pkgcontext "github.com/arymaulanamalik/ai-engineer-playground/ai-playground/pkg/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Setup(conf *configs.Config) {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(parseLevel(conf.LogLevel)),
		Encoding:    "json",
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{
			"stderr",
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "time",
			LevelKey:   "level",
			MessageKey: "message",
			CallerKey:  "caller",

			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	log, err = cfg.Build(
		zap.AddCaller(),
	)

	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(log)

	_ = os.Stdout.Sync()
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return zap.DebugLevel
	case "WARN":
		return zap.WarnLevel
	case "ERROR":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func withCtx(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return log
	}

	reqID := pkgcontext.RequestID(ctx)

	return log.With(
		zap.String("req_id", reqID),
	)
}

func Debug(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	withCtx(ctx).Debug(msg, fields...)
}

func Info(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	withCtx(ctx).Info(msg, fields...)
}

func Warn(
	ctx context.Context,
	msg string,
	fields ...zap.Field,
) {
	withCtx(ctx).Warn(msg, fields...)
}

func Error(
	ctx context.Context,
	err error,
	msg string,
	fields ...zap.Field,
) {
	fields = append(
		fields,
		zap.Error(err),
	)

	withCtx(ctx).Error(msg, fields...)
}
