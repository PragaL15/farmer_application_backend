package main

// @title           Go New Backend API
// @version         1.0
// @description     This is the API documentation for Go Fiber App.
// @termsOfService  http://yourterms.com
// @contact.name    Pragalya Kanakaraj
// @contact.email   your-email@example.com
// @host            localhost:3000
// @BasePath        /
// @schemes         http

import (
	"farmerapp/go_backend/db"
	"farmerapp/routes"
	"farmerapp/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	logFile := utils.InitLogger()
	defer logFile.Close()
	utils.Logger.Println("Server is starting...")

	db.ConnectDB()
	defer db.CloseDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8100, http://localhost:5500",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Request logger middleware
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		utils.Logger.Printf(
			"METHOD=%s PATH=%s STATUS=%d DURATION=%v\n",
			c.Method(), c.Path(), c.Response().StatusCode(), time.Since(start),
		)
		return err
	})

	// Swagger documentation route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Sample GET endpoint to test CORS
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Register your API routes
	routes.RegisterRoutes(app)

	// Graceful shutdown
	go func() {
		utils.Logger.Printf("Server is running on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			utils.Logger.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	utils.Logger.Println("Shutting down server gracefully...")
	if err := app.Shutdown(); err != nil {
		utils.Logger.Fatalf("Server forced shutdown: %v", err)
	}
	utils.Logger.Println("Server stopped cleanly")
}
