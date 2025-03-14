package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/handlers"
  Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
)

func RegisterRoutes(app *fiber.App) {
	// GET Routes
	app.Get("/getUsers", Masterhandlers.GetAllUsers)
	app.Get("/getCategories", Masterhandlers.GetCategories)
	app.Get("/getDrivers", Masterhandlers.GetDrivers)
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

	app.Get("/products/:id", Masterhandlers.GetProductByID)
  app.Get("/categories/:category_id", Masterhandlers.GetCategoryByID)
  app.Get("/allOrderDetails/:id", handlers.GetOrderDetailsByOrderID)
  app.Get("/getProductByCatId/:category_id", Masterhandlers.GetProductsByCategoryID)

  
	// app.Get("/getOrders", handlers.GetOrders)
	app.Get("/getModeOfPayments", handlers.GetModeOfPayments)
	app.Get("/getListPaymentMethod", Masterhandlers.GetListPaymentMethods)
	app.Get("/getCompleteOrderDetails", handlers.GetOrderDetails)
	app.Get("/getOrderHistoryDetails", handlers.GetOrderHistory)
	app.Get("/getInvoiceDetails", handlers.GetInvoiceDetails)
	app.Get("/getDailyPriceDetails", handlers.GetDailyPriceUpdates)

	// POST Routes
	app.Post("/user-bank-details", handlers.InsertUserBankDetail)
	app.Post("/categoryDetails", Masterhandlers.InsertCategory)
	app.Post("/driverDetails", Masterhandlers.InsertDriver)
	app.Post("/locationDetails", Masterhandlers.InsertLocation)
	app.Post("/mandiDetails", Masterhandlers.InsertMasterMandi)
	app.Post("/productDetails", Masterhandlers.InsertMasterProduct)
	app.Post("/stateDetails", Masterhandlers.InsertMasterState)
	app.Post("/vehicleDetails", Masterhandlers.InsertMasterVehicle)
	app.Post("/violationDetails", Masterhandlers.InsertMasterViolation)
	app.Post("/userTableDetails", Masterhandlers.InsertUser)
	app.Post("/businessDetails", Masterhandlers.InsertBusiness)
	app.Post("/orderStatusDetails", Masterhandlers.InsertOrderStatus)
	app.Post("/ordersDetails", handlers.InsertOrder)

	// PUT Routes (Updating)
	app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
	app.Put("/categoryUpdate", Masterhandlers.UpdateCategory)
	app.Put("/driverUpdate", Masterhandlers.UpdateDriver)
	app.Put("/locationUpdate", Masterhandlers.UpdateLocation)
	app.Put("/mandiUpdate", Masterhandlers.UpdateMasterMandi)
	app.Put("/productUpdate", Masterhandlers.UpdateMasterProduct)
	app.Put("/statesUpdate", Masterhandlers.UpdateMasterState)
	app.Put("/vehicleUpdate", Masterhandlers.UpdateMasterVehicle)
	app.Put("/violationUpdate", Masterhandlers.UpdateMasterViolation)
	app.Put("/usertableUpdate", Masterhandlers.UpdateUser)
	app.Put("/businessUpdate", Masterhandlers.UpdateBusiness)
	app.Put("/orderStatusUpdate", Masterhandlers.UpdateOrderStatus)
	app.Put("/ordersUpdate", handlers.UpdateOrder)
	app.Put("/dailyPriceUpdate", handlers.UpdateDailyPrice)
}