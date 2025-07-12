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
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "farmerapp/docs"
	"farmerapp/go_backend/db"
	"farmerapp/routes"
	"farmerapp/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title           My API
// @version         1.0
// @description     This is the API documentation for Go Fiber App.
// @host            localhost:3000
// @BasePath        /
// @schemes         http
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
	app.Use(cors.New())

	/*
		app.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowCredentials: true,
		})) */

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

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Your API routes
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
