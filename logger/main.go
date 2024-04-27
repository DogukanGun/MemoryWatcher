package logger

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

type Log struct {
	Logger *zap.Logger
}

type LogErrorStruct struct {
	Message string `json:"message"`
}
type LogInfoStruct struct {
	Message string `json:"message"`
}

func Init() *Log {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	return &Log{Logger: logger}
}

func LogError(errorStruct LogErrorStruct) {
	fmt.Printf("Error: %s", errorStruct.Message)
	os.Exit(1)
}

func LogInfo(infoStruct LogInfoStruct) {
	fmt.Printf("INFO: [%s]\n", infoStruct.Message)
}
