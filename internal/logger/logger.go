package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func InitLogger(logFilePath string) error {
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logger = log.New(logFile, "", log.LstdFlags)
	logger.Println("Логгер успешно инициализирован")
	return nil
}

func GetLogger() *log.Logger {
	if logger == nil {
		log.Fatalf("Логгер не инициализирован. Вызовите InitLogger() перед использованием.")
	}
	return logger
}
