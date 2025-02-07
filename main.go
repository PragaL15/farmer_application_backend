package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" 
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/PragaL15/go_newBackend/handlers"
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

	logger := log.New(logFile, "CUSTOM: ", log.Ldate|log.Ltime|log.Lshortfile)

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
	logger.Println(" Server is starting...")

	db.ConnectDB()
	defer db.CloseDB()

	app := fiber.New()

	app.Get("/users", handlers.GetAllUsers)

	app.Post("/user-bank-details", handlers.InsertUserBankDetail)
	app.Post("/categoryDetails",handlers.InsertCategory)
	app.Post("/driverDetails",handlers.InsertDriver)
	app.Post("/locationDetails",handlers.InsertLocation)
	app.Post("/mandiDetails",handlers.InsertMasterMandi)
	app.Post("/productDetails",handlers.InsertMasterProduct)
	app.Post("/stateDetails",handlers.InsertMasterState)
	app.Post("/vehicleDetails",handlers.InsertMasterVehicle)
	app.Post("/violationDetails",handlers.InsertMasterViolation)
	app.Post("/userTableDetails",handlers.InsertUser)

	app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
	app.Put("/categoryUpdate", handlers.UpdateCategory)
	app.Put("/driverUpdate", handlers.UpdateDriver)
	app.Put("/locationUpdate", handlers.UpdateLocation)
	app.Put("/mandiUpdate", handlers.UpdateMasterMandi)
	app.Put("/productUpdate", handlers.UpdateMasterProduct)
	app.Put("/statesUpdate", handlers.UpdateMasterState)
	app.Put("/vehicleUpdate", handlers.UpdateMasterVehicle)
	app.Put("/violationUpdate", handlers.UpdateMasterViolation)
	app.Put("/usertableUpdate", handlers.UpdateUser)

	app.Delete("/user-bank-details/:id", handlers.DeleteUserBankDetail)
	app.Delete("/categoryDelete/:id", handlers.DeleteCategory)
	app.Delete("/driverDelete/:id", handlers.DeleteDriver)
	app.Delete("/locationDelete/:id", handlers.DeleteLocation)
	app.Delete("/mandiDelete/:id", handlers.DeleteMasterMandi)
	app.Delete("/productDelete/:id", handlers.DeleteMasterProduct)
	app.Delete("/stateDelete/:id", handlers.DeleteMasterState)
	app.Delete("/vehicleDelete/:id", handlers.DeleteMasterVehicle)
	app.Delete("/violationDelete/:id", handlers.DeleteMasterViolation)
	app.Delete("/usertableDelete/:id", handlers.DeleteUser)

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
