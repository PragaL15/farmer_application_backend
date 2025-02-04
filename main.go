package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/PragaL15/go_newBackend/handlers"
)

func main() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	log.SetOutput(logFile)
	defer logFile.Close()

	db.ConnectDB()
	defer db.CloseDB()

	app := fiber.New()

	app.Get("/users", handlers.GetAllUsers)
	app.Post("/user-bank-details", handlers.InsertUserBankDetail)
	app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
	app.Delete("/user-bank-details/:id", handlers.DeleteUserBankDetail)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()
    
	log.Println("Server is running on port 3000")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server gracefully...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}
	log.Println("Server stopped cleanly")
}
