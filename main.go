package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/PragaL15/go_newBackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func createLogger() (*log.Logger, *os.File) {
	currentDate := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("logs/%s.log", currentDate)

	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatalf("Error creating logs directory: %v", err)
	}

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	logger := log.New(logFile, "SERVER: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger, logFile
}

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	logger, logFile := createLogger()
	defer logFile.Close()
	logger.Println("Server is starting...")
	db.ConnectDB()
	defer db.CloseDB()
	app := fiber.New()
	routes.RegisterRoutes(app)
	go func() {
		logger.Printf("Server is running on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			logger.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logger.Println("Shutting down server gracefully...")
	if err := app.Shutdown(); err != nil {
		logger.Fatalf("Server forced shutdown: %v", err)
	}
	logger.Println("Server stopped cleanly")
}
