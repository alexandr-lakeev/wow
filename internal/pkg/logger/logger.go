package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logg *zap.Logger
}

func New() (*Logger, error) {
	logg, err := initLogger()
	if err != nil {
		return nil, err
	}

	return &Logger{
		logg: logg,
	}, nil
}

func (l Logger) Debug(msg string) {
	l.logg.Debug(msg)
}

func (l Logger) Info(msg string) {
	l.logg.Info(msg)
}

func (l Logger) Warning(msg string) {
	l.logg.Warn(msg)
}

func (l Logger) Error(msg string) {
	l.logg.Error(msg)
}

func (l Logger) Panic(msg string) {
	l.logg.Panic(msg)
}

func initLogger() (*zap.Logger, error) {
	atom := zap.NewAtomicLevel()

	var logger *zap.Logger

	logger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atom,
	))

	atom.SetLevel(zap.DebugLevel)

	return logger, nil
}
