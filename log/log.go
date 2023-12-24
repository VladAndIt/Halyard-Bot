package log

import (
	"context"
	"log/slog"
	"os"
)

type ctxLogger struct{}

func ContextWithLogger(ctx context.Context) context.Context {
	var log *slog.Logger

	switch profile {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return context.WithValue(ctx, ctxLogger{}, log)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(ctxLogger{}).(*slog.Logger); ok {
		return logger
	}
	return nil
}
