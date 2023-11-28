package log

import "go.uber.org/zap"

type Logger interface {
	// Debug uses fmt.Sprint to construct and log a message at DEBUG level
	Debug(args ...interface{})
	// Info uses fmt.Sprint to construct and log a message at INFO level
	Info(args ...interface{})
	// Error uses fmt.Sprint to construct and log a message at ERROR level
	Error(args ...interface{})

	// Debugf uses fmt.Sprintf to construct and log a message at DEBUG level
	Debugf(format string, args ...interface{})
	// Infof uses fmt.Sprintf to construct and log a message at INFO level
	Infof(format string, args ...interface{})
	// Errorf uses fmt.Sprintf to construct and log a message at ERROR level
	Errorf(format string, args ...interface{})
}

type logger struct {
	*zap.SugaredLogger
}

func New() Logger {
	l, _ := zap.NewProduction()
	return NewWithZap(l)
}

func NewWithZap(l *zap.Logger) Logger {
	return &logger{l.Sugar()}
}
