package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zlogger *zap.SugaredLogger
)

func init() {
	cores := []zapcore.Core{
		zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	}

	core := zapcore.NewTee(cores...)
	l := zap.New(core)
	zlogger = l.Sugar().WithOptions(zap.AddCallerSkip(1))
}

func Debug(format string, v ...interface{}) {
	zlogger.Debugf(format, v...)
}

func Info(format string, v ...interface{}) {
	zlogger.Infof(format, v...)
}

func Warn(format string, v ...interface{}) {
	zlogger.Warnf(format, v...)
}

func Error(format string, v ...interface{}) {
	zlogger.Errorf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	zlogger.Fatalf(format, v...)
}
