package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/PragaL15/go_newBackend/routes"
	"github.com/PragaL15/go_newBackend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize logger from utils
	logFile := utils.InitLogger()
	defer logFile.Close()
	utils.Logger.Println("Server is starting...")

	// Connect to DB
	db.ConnectDB()
	defer db.CloseDB()

	app := fiber.New()

	// Middleware: Log all requests
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		utils.Logger.Printf(
			"METHOD=%s PATH=%s STATUS=%d DURATION=%v\n",
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			time.Since(start),
		)
		return err
	})

	// Register all routes
	routes.RegisterRoutes(app)

	// Start server in goroutine
	go func() {
		utils.Logger.Printf("Server is running on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			utils.Logger.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	utils.Logger.Println("Shutting down server gracefully...")
	if err := app.Shutdown(); err != nil {
		utils.Logger.Fatalf("Server forced shutdown: %v", err)
	}
	utils.Logger.Println("Server stopped cleanly")
}
