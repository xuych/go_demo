package util

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/xerrors"
)

var (
	Logger *zap.Logger
)

func ZapLogInit(level *string) error {
	if level == nil {
		return xerrors.New("not set Log level")
	}
	var logLevel zapcore.Level
	lowLevel := strings.ToLower(*level)
	switch lowLevel {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		return xerrors.New("un support log level: " + lowLevel)
	}
	core := zapcore.NewCore(getEncoder(), getLogWriter(), logLevel)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger
	return nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}
