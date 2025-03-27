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
	app.Get("/getMandiDetails", Masterhandlers.GetAllMandiDetails)
	app.Get("/getAllLanguages", Masterhandlers.GetAllLanguages)
	app.Get("/getProducts", Masterhandlers.GetAllProducts)
	app.Get("/getStates", Masterhandlers.GetStates)
	app.Get("/getVehicles", Masterhandlers.GetVehicles)
	app.Get("/getViolations", Masterhandlers.GetViolations)
	app.Get("/getBusinesses", Masterhandlers.GetAllBusinesses)
	app.Get("/getDriverViolation", handlers.GetDriverViolations)
	app.Get("/getOrderStatus", Masterhandlers.GetOrderStatuses)
	app.Get("/getBusinessTypes", Masterhandlers.GetBusinessTypes)
	app.Get("/getModeOfPayments", handlers.GetModeOfPayments)
	app.Get("/getListPaymentMethod", Masterhandlers.GetAllCashPaymentsType)
	app.Get("/getCompleteOrderDetails", handlers.GetOrderDetails)
	app.Get("/getOrderHistoryDetails", handlers.GetOrderHistory)
	app.Get("/getInvoiceDetails", handlers.GetInvoiceDetails)
	app.Get("/getDailyPriceDetails", handlers.GetDailyPriceUpdates)
	app.Get("/getPaymentMode", handlers.GetPaymentModes)
	app.Get("/getBusinessCategory", Masterhandlers.GetBusinessCategories)
	app.Get("/getBusinessBranches", Masterhandlers.GetAllBusinessBranches)
  app.Get("/getBusinessUsers", Masterhandlers.GetAllBusinessUsers)
  app.Get("/getProductCategoryRegionalName", Masterhandlers.GetProductCategoryRegionalNames)

	app.Get("/products/:id", Masterhandlers.GetProductByID)
	app.Get("/getBusinesses/:id", Masterhandlers.GetBusinessByID)
	app.Get("/getMandiDetails/:id", Masterhandlers.GetMandiDetailsByID)
	app.Get("/getBusinessesbranch/:id", Masterhandlers.GetBusinessBranchByID)
  app.Get("/categories/:category_id", Masterhandlers.GetCategoryByID)
  app.Get("/allOrderDetails/:id", handlers.GetOrderDetailsByOrderID)
  app.Get("/getProductByCatId/:category_id", Masterhandlers.GetProductsByCategoryID)
  app.Get("/getPaymentModeById/:id", handlers.GetPaymentModeByID)
  app.Get("/getOrderStatusById/:id", Masterhandlers.GetOrderStatusByID)
  app.Get("/getPaymentTypeById/:id", Masterhandlers.GetCashPaymentByID)
  app.Get("/getBusinessTypeById/:id", Masterhandlers.GetBusinessTypeByID)
  app.Get("/getDriverById/:id", Masterhandlers.GetDriverByID)
  app.Get("/getDriverById/:id", Masterhandlers.GetBusinessBranchByID)
	app.Get("/getBusinessUser/:id", Masterhandlers.GetBusinessUserByID) 
	app.Get("/getProductCategoryRegional/:id", Masterhandlers.GetProductCategoryRegionalNameByID) 

	// POST Routes
	app.Post("/user-bank-details", handlers.InsertUserBankDetail)
	app.Post("/categoryDetails", Masterhandlers.InsertCategory)
	app.Post("/driverDetails", Masterhandlers.InsertDriver)
	app.Post("/locationDetails", Masterhandlers.InsertLocation)
	app.Post("/mandiDetails", Masterhandlers.InsertMandiDetails)
	app.Post("/productDetails", Masterhandlers.InsertProduct)
	app.Post("/stateDetails", Masterhandlers.InsertMasterState)
	app.Post("/vehicleDetails", Masterhandlers.InsertMasterVehicle)
	app.Post("/violationDetails", Masterhandlers.InsertMasterViolation)
	app.Post("/userTableDetails", Masterhandlers.InsertUser)
	app.Post("/businessDetails", Masterhandlers.InsertBusiness)
	app.Post("/orderStatusDetails", Masterhandlers.InsertOrderStatus)
	app.Post("/ordersDetails", handlers.InsertOrder)
	app.Post("/paymentModes", handlers.InsertPaymentMode)
	app.Post("/InsertLanguages", Masterhandlers.InsertLanguage)
	app.Post("/businessTypeDetails", Masterhandlers.InsertBusinessType)
	app.Post("/listPaymentsDetails", Masterhandlers.InsertPaymentType)
	app.Post("/InsertBusinessCategory", Masterhandlers.InsertBusinessCategory)
	app.Post("/business-branches", Masterhandlers.InsertBusinessBranch)
	app.Post("/PostProductCategoryRegional", Masterhandlers.InsertProductCategoryRegionalName)

	// PUT Routes (Updating)
	app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
	app.Put("/categoryUpdate", Masterhandlers.UpdateCategory)
	app.Put("/driverUpdate", Masterhandlers.UpdateDriver)
	app.Put("/locationUpdate", Masterhandlers.UpdateLocation)
	app.Put("/mandiUpdate", Masterhandlers.UpdateMandiDetails)
	app.Put("/languagesUpdate", Masterhandlers.UpdateLanguage)
	app.Put("/productUpdate", Masterhandlers.UpdateProductHandler)
	app.Put("/statesUpdate", Masterhandlers.UpdateMasterState)
	app.Put("/vehicleUpdate", Masterhandlers.UpdateMasterVehicle)
	app.Put("/violationUpdate", Masterhandlers.UpdateMasterViolation)
	app.Put("/usertableUpdate", Masterhandlers.UpdateUser)
	app.Put("/businessUpdate", Masterhandlers.UpdateBusiness)
	app.Put("/orderStatusUpdate", Masterhandlers.UpdateOrderStatus)
	app.Put("/ordersUpdate", handlers.UpdateOrder)
	app.Put("/dailyPriceUpdate", handlers.UpdateDailyPrice)
	app.Put("/paymentModeUpdate", handlers.UpdatePaymentMode)
	app.Put("/businessTypeUpdate", Masterhandlers.UpdateBusinessType)
	app.Put("/PaymentTypeUpdate", Masterhandlers.UpdatePaymentType)
	app.Put("/businessCategoryUpdate", Masterhandlers.UpdateBusinessCategory)
	app.Put("/branch/:id", Masterhandlers.UpdateBusinessBranch)
	app.Put("/businessUserDetailesUpdate", Masterhandlers.UpdateBusinessUser) 
	app.Put("/productCategoryRegionalUpdate", Masterhandlers.UpdateProductCategoryRegionalName) 

}

