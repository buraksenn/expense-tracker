package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zlogger *zap.SugaredLogger
)

func Init(f *os.File) {
	cores := []zapcore.Core{
		zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(f), zapcore.DebugLevel),
	}

	if f != os.Stdout {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(f), zapcore.InfoLevel))
	}

	core := zapcore.NewTee(cores...)
	l := zap.New(core)
	zlogger = l.Sugar()
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

func Panic(format string, v ...interface{}) {
	zlogger.Panicf(format, v...)
}
