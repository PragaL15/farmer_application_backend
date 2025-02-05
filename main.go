package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" // Load .env file
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/PragaL15/go_newBackend/handlers"
)

// Create a logger that writes logs to a daily log file
func createLogger() (*log.Logger, *os.File) {
	// Get current date in YYYY-MM-DD format
	currentDate := time.Now().Format("2006-01-02")

	// Create a filename using the current date
	logFileName := fmt.Sprintf("logs/%s.log", currentDate)

	// Ensure the "logs" directory exists
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatalf("Error creating logs directory: %v", err)
	}

	// Open (or create) the log file in append mode
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Create a custom logger with date, time, and source file info
	logger := log.New(logFile, "CUSTOM: ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger, logFile
}

func main() {
	// Load environment variables from .env file
	_ = godotenv.Load() // Optional: If .env is missing, it won’t throw an error

	// Get PORT from environment variable, default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}

	// Initialize logger
	logger, logFile := createLogger()
	defer logFile.Close() // Ensure the file is closed when the program exits

	logger.Println("🚀 Server is starting...")

	// Initialize database connection
	db.ConnectDB()
	defer db.CloseDB()

	// Initialize Fiber app
	app := fiber.New()

	// Define routes
	app.Get("/users", handlers.GetAllUsers)
	app.Post("/user-bank-details", handlers.InsertUserBankDetail)
	app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
	app.Delete("/user-bank-details/:id", handlers.DeleteUserBankDetail)

	// Run server in a goroutine
	go func() {
		logger.Printf("✅ Server is running on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			logger.Fatalf("❌ Server error: %v", err)
		}
	}()

	// Handle OS signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Println("🛑 Shutting down server gracefully...")
	if err := app.Shutdown(); err != nil {
		logger.Fatalf("❌ Server forced shutdown: %v", err)
	}

	logger.Println("✅ Server stopped cleanly")
}
