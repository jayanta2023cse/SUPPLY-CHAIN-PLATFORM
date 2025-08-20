// logger/logger.go
package logger

import (
	"log"
	"os"
	"sync"
)

var (
	logger   *log.Logger
	initOnce sync.Once
)

func InitLogger() *log.Logger {
	initOnce.Do(func() {
		logFile, err := os.OpenFile("supply_chain_platform.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		logger = log.New(logFile, "", log.LstdFlags)
	})
	return logger
}

func GetLogger(service string) *log.Logger {
	if logger == nil {
		InitLogger()
	}
	// Return logger with service-specific prefix
	return log.New(logger.Writer(), "["+service+"] ", log.LstdFlags)
}
