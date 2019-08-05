package tools

import "go.uber.org/zap"

func GetLogger() (sLogger *zap.SugaredLogger) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sLogger = logger.Sugar()

	return sLogger
}