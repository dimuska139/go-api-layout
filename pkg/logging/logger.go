package logging

import (
	"context"
	"log/slog"
	"os"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

var logger *slog.Logger

func getSlogLevel(level LogLevel) slog.Level {
	switch level {
	case LogLevelDebug:
		return slog.LevelDebug

	case LogLevelInfo:
		return slog.LevelInfo

	case LogLevelWarn:
		return slog.LevelWarn

	case LogLevelError:
		return slog.LevelError

	default:
		return slog.LevelInfo
	}
}

func InitLogger(loglevel LogLevel) {
	var opts *slog.HandlerOptions
	if loglevel != "" {
		opts = &slog.HandlerOptions{
			Level: getSlogLevel(loglevel),
		}
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger = slog.New(handler)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	if logger == nil {
		return
	}

	Error(ctx, msg, args...)

	os.Exit(1)
}

func Write(ctx context.Context, level LogLevel, msg string, args ...any) {
	if logger == nil {
		return
	}

	logger.Log(ctx, getSlogLevel(level), msg, args...)
}

func Debug(_ context.Context, msg string, args ...any) {
	if logger == nil {
		return
	}

	logger.Debug(msg, args...)
}

func Info(_ context.Context, msg string, args ...any) {
	if logger == nil {
		return
	}

	logger.Info(msg, args...)
}

func Warn(_ context.Context, msg string, args ...any) {
	if logger == nil {
		return
	}

	logger.Warn(msg, args...)
}

func Error(_ context.Context, msg string, args ...any) {
	if logger == nil {
		return
	}

	logger.Error(msg, args...)
}
