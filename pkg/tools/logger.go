// Package tools provides various utilities which can be useful for developers.
package tools

import "go.uber.org/zap"

// Returns the logger to be used for logging purposes.
func GetLogger() (sLogger *zap.SugaredLogger) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sLogger = logger.Sugar()

	return sLogger
}
