package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/PragaL15/go_newBackend/handlers"
	"github.com/PragaL15/go_newBackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	// routes "github.com/PragaL15/go_newBackend/routes"
	Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
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

	//get API's
	app.Get("/getUsers", Masterhandlers.GetAllUsers)
	app.Get("/getCategories", Masterhandlers.GetCategories)
	app.Get("/getDrivers",Masterhandlers.GetDrivers)
	app.Get("/getLocations", Masterhandlers.GetLocations)
	app.Get("/getMandis", Masterhandlers.GetMandi)
	app.Get("/getProducts", Masterhandlers.GetProducts) 
	app.Get("/getStates", Masterhandlers.GetStates)
	app.Get("/getVehicles", Masterhandlers.GetVehicles)
	app.Get("/getViolations", Masterhandlers.GetViolations)  
	app.Get("/getBusinesses", Masterhandlers.GetBusinesses)  
	app.Get("/getDriverViolation", handlers.GetDriverViolations)
	app.Get("/getOrderStatus", Masterhandlers.GetOrderStatuses)
	app.Get("/getBusinessStatus", handlers.GetBusinessTypes)
//app.Get("/getOrders", handlers.GetOrders)
	app.Get("/getModeOfPayments", handlers.GetModeOfPayments)
	app.Get("/getListPaymentMethod", Masterhandlers.GetListPaymentMethods)
	app.Get("/getCompleteOrderDetails", handlers.GetOrderDetails)
	app.Get("/getOrderHistoryDetails", handlers. GetOrderHistory)
	app.Get("/getInvoiceDetails", handlers.GetInvoiceDetails)

	//posting API's
	app.Post("/user-bank-details", handlers.InsertUserBankDetail)
	app.Post("/categoryDetails",Masterhandlers.InsertCategory)
	app.Post("/driverDetails",Masterhandlers.InsertDriver)
	app.Post("/locationDetails",Masterhandlers.InsertLocation)
	app.Post("/mandiDetails",Masterhandlers.InsertMasterMandi)
	app.Post("/productDetails",Masterhandlers.InsertMasterProduct)
	app.Post("/stateDetails",Masterhandlers.InsertMasterState)
	app.Post("/vehicleDetails",Masterhandlers.InsertMasterVehicle)
	app.Post("/violationDetails",Masterhandlers.InsertMasterViolation)
	app.Post("/userTableDetails",Masterhandlers.InsertUser)
	app.Post("/businessDetails",Masterhandlers.InsertBusiness)

	app.Post("/orderStatusDetails",Masterhandlers.InsertOrderStatus)
	app.Post("/ordersDetails",handlers.InsertOrder)

	//Updating API's
	app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
	app.Put("/categoryUpdate", Masterhandlers.UpdateCategory)
	app.Put("/driverUpdate", Masterhandlers.UpdateDriver)
	app.Put("/locationUpdate", Masterhandlers.UpdateLocation)
	app.Put("/mandiUpdate", Masterhandlers.UpdateMasterMandi)
	app.Put("/productUpdate",Masterhandlers.UpdateMasterProduct)
	app.Put("/statesUpdate", Masterhandlers.UpdateMasterState)
	app.Put("/vehicleUpdate", Masterhandlers.UpdateMasterVehicle)
	app.Put("/violationUpdate", Masterhandlers.UpdateMasterViolation)
	app.Put("/usertableUpdate", Masterhandlers.UpdateUser)
	app.Put("/businessUpdate", Masterhandlers.UpdateBusiness)

	app.Put("/orderStatusUpdate", Masterhandlers.UpdateOrderStatus)
	app.Put("/ordersUpdate", handlers.UpdateOrder)


	//Deleting API's
	// app.Delete("/user-bank-details/:id", handlers.DeleteUserBankDetail)
	// app.Delete("/categoryDelete/:id", handlers.DeleteCategory)
	// app.Delete("/driverDelete/:id", handlers.DeleteDriver)
	// app.Delete("/locationDelete/:id", handlers.DeleteLocation)
	// app.Delete("/mandiDelete/:id", handlers.DeleteMasterMandi)
	// app.Delete("/productDelete/:id", handlers.DeleteMasterProduct)
	// app.Delete("/stateDelete/:id", handlers.DeleteMasterState)
	// app.Delete("/vehicleDelete/:id", handlers.DeleteMasterVehicle)
	// app.Delete("/violationDelete/:id", handlers.DeleteMasterViolation)
	// app.Delete("/usertableDelete/:id", handlers.DeleteUser)
	// app.Delete("/businessDelete/:id", handlers.DeleteBusiness)

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




