package logger

import (
	"context"

	"go.uber.org/zap"
)

func withContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return zlogger
	}

	chatID, ok := ctx.Value("chatID").(string)
	if !ok {
		return zlogger
	}
	return zlogger.With("chat_id:", chatID)
}

func DebugC(ctx context.Context, format string, v ...interface{}) {
	withContext(ctx).Debugf(format, v...)
}

func InfoC(ctx context.Context, format string, v ...interface{}) {
	withContext(ctx).Infof(format, v...)
}

func ErrorC(ctx context.Context, format string, v ...interface{}) {
	withContext(ctx).Errorf(format, v...)
}
