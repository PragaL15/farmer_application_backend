package utils
import (
	"fmt"
	"log"
	"os"
	"time"
)

var Logger *log.Logger

func InitLogger() *os.File {
	currentDate := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("logs/%s.log", currentDate)
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatalf("Error creating logs directory: %v", err)
	}
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	Logger = log.New(logFile, "API: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logFile
}