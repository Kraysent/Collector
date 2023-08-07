package log

import (
	"go.uber.org/zap"
)

var globalLogger *zap.Logger

func InitLogger(stdoutPath, stderrPath string) error {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.Development = true
	config.OutputPaths = []string{"stdout"}
	if stdoutPath != "" {
		config.OutputPaths = append(config.OutputPaths, stdoutPath)
	}
	config.ErrorOutputPaths = []string{"stderr"}
	if stderrPath != "" {
		config.ErrorOutputPaths = append(config.ErrorOutputPaths, stderrPath)
	}

	logger, err := config.Build()
	if err != nil {
		return err
	}

	globalLogger = logger

	return nil
}

func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	globalLogger.Fatal(msg, fields...)
}
