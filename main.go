package main

import (
        "os"
        "os/signal"
        "path/filepath"
        "syscall"
        "time"

        "github.com/gofiber/fiber/v2"
        "github.com/gofiber/fiber/v2/middleware/cors"
        "github.com/gofiber/swagger"
        "github.com/joho/godotenv"

<<<<<<< HEAD
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
=======
        _ "farmerapp/docs"
        "farmerapp/go_backend/db"
        "farmerapp/routes"
        "farmerapp/utils"
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796
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

<<<<<<< HEAD
	app := fiber.New()
	app.Use(cors.New())

	/*
		app.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowCredentials: true,
		})) */
=======
        db.ConnectDB()
        defer db.CloseDB()
>>>>>>> 89c8837d6c691851780de9cd6d15506fa99fd796

        app := fiber.New()

        // Enable CORS
        app.Use(cors.New(cors.Config{
                AllowOrigins: "http://localhost:3000, http://api.ekadyu.org:8081, http://localhost:8100, http://localhost:5500",
                AllowHeaders: "Origin, Content-Type, Accept, Authorization",
                AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
        }))

        // Serve static files
        wd, _ := os.Getwd()
        staticPath := filepath.Join(wd, "farmer_application_backend-master", "static")
        app.Static("/", staticPath)

 // Serve index.html at root
        app.Get("/", func(c *fiber.Ctx) error {
                return c.SendFile(filepath.Join(staticPath, "index.html"))
        })

        // Request logging
        app.Use(func(c *fiber.Ctx) error {
                start := time.Now()
                err := c.Next()
                utils.Logger.Printf(
                        "METHOD=%s PATH=%s STATUS=%d DURATION=%v\n",
                        c.Method(), c.Path(), c.Response().StatusCode(), time.Since(start),
                )
                return err
        })

        // Swagger docs
        app.Get("/swagger/*", swagger.HandlerDefault)

        // Register routes
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
