package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/PragaL15/go_newBackend/docs"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/PragaL15/go_newBackend/routes"
	"github.com/PragaL15/go_newBackend/utils"
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
