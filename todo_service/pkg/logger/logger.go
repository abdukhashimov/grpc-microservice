package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field

var (
	Int    = zap.Int
	String = zap.String
	Error  = zap.Error
	Bool   = zap.Bool
	Any    = zap.Any
)

// Logger ...
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type loggerImpl struct {
	zap *zap.Logger
}

var (
	customTimeFormat string
)

// New ...
func New(level string, namespace string) Logger {
	if level == "" {
		// when we do not send the level it is `info` by default
		// LevelInfo is defined in const.go
		level = LevelInfo
	}

	logger := loggerImpl{
		zap: newZapLogger(level, time.RFC3339),
	}

	logger.zap = logger.zap.Named(namespace)

	zap.RedirectStdLog(logger.zap)

	return &logger
}

func (l *loggerImpl) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, fields)
}

func (l *loggerImpl) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l *loggerImpl) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}

func (l *loggerImpl) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}

func (l *loggerImpl) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, fields...)
}
