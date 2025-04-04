package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/handlers"
  Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
)

func RegisterRoutes(app *fiber.App) {

	// GET Routes
	app.Get("/getUsers", Masterhandlers.GetAllUsers)
	//app.Get("/getCategories", Masterhandlers.GetCategories)
	app.Get("/getDrivers", Masterhandlers.GetDrivers)
	app.Get("/getLocations", Masterhandlers.GetLocations)
	app.Get("/getMandiDetails", Masterhandlers.GetAllMandiDetails)
	app.Get("/getAllLanguages", Masterhandlers.GetAllLanguages)
	app.Get("/getProducts", Masterhandlers.GetAllProducts)
	app.Get("/getStates", Masterhandlers.GetAllStates)
	app.Get("/getVehicles", Masterhandlers.GetAllVehicles)
	app.Get("/getViolations", Masterhandlers.GetViolations)
	app.Get("/getBusinesses", Masterhandlers.GetAllBusinesses)
	app.Get("/getDriverViolation", handlers.GetDriverViolations)
	app.Get("/getOrderStatus", Masterhandlers.GetAllOrderStatuses)
	app.Get("/getBusinessTypes", Masterhandlers.GetBusinessTypes)
	app.Get("/getModeOfPayments", Masterhandlers.GetPaymentModes)
	app.Get("/getListPaymentMethod", Masterhandlers.GetAllCashPaymentsType)
	app.Get("/getOrderHistoryDetails", handlers.GetOrderHistory)
	//app.Get("/getInvoiceDetails", handlers.GetInvoiceDetails)
	app.Get("/getDailyPriceDetails", handlers.GetPriceHandler)
	app.Get("/getPaymentMode", Masterhandlers.GetPaymentModes)
	app.Get("/getBusinessCategory", Masterhandlers.GetBusinessCategories)
	app.Get("/getBusinessBranches", Masterhandlers.GetAllBusinessBranches)
	app.Get("/getAllOrderDetails1", handlers.GetOrderDetailsHandler)
  app.Get("/getBusinessUsers", Masterhandlers.GetAllBusinessUsers)
  app.Get("/getProductCategoryRegionalName", Masterhandlers.GetProductCategoryRegional)

	// GET each by id's
	app.Get("/getProducts/:id", Masterhandlers.GetProductByID)
	app.Get("/getUsers/:id", Masterhandlers.GetUserByID)
	app.Get("/getBusinesses/:id", Masterhandlers.GetBusinessByID)
	app.Get("/getMandiDetails/:id", Masterhandlers.GetMandiDetailsByID)
  app.Get("/getCategories/:category_id", Masterhandlers.GetCategoryByID)
  app.Get("/getStates/:state_id", Masterhandlers.GetStateByID)
  app.Get("/getProductByCatId/:category_id", Masterhandlers.GetProductsByCategoryID)
  app.Get("/getPaymentModeById/:id", Masterhandlers.GetPaymentModeByID)
  app.Get("/getOrderStatusById/:id", Masterhandlers.GetOrderStatusByID)
  app.Get("/getPaymentTypeById/:id", Masterhandlers.GetCashPaymentByID)
  app.Get("/getBusinessTypeById/:id", Masterhandlers.GetBusinessTypeByID)
  app.Get("/getDriverById/:id", Masterhandlers.GetDriverByID)
  app.Get("/getBusinessesbranch/:id", Masterhandlers.GetBusinessBranchByID)
	app.Get("/getBusinessUser/:id", Masterhandlers.GetBusinessUserByID) 
	app.Get("/getProductCategoryRegional/:id", Masterhandlers.GetProductCategoryRegionalByID) 
  app.Get("/invoice/:invoice_id", handlers.GetInvoiceDetails)

	// POST Routes
	app.Post("/user-bank-details", handlers.InsertUserBankDetail)
	app.Post("/categoryDetails", Masterhandlers.InsertCategory)
	app.Post("/driverDetails", Masterhandlers.InsertDriver)
	app.Post("/locationDetails", Masterhandlers.InsertLocation)
	app.Post("/mandiDetails", Masterhandlers.InsertMandiDetails)
	app.Post("/productDetails", Masterhandlers.InsertProduct)
	app.Post("/stateDetails", Masterhandlers.InsertState)
	app.Post("/vehicleDetails", Masterhandlers.InsertVehicle)
	app.Post("/violationDetails", Masterhandlers.InsertMasterViolation)
	app.Post("/userTableDetails", Masterhandlers.InsertUser)
	app.Post("/businessDetails", Masterhandlers.InsertBusiness)
	app.Post("/orderStatusDetails", Masterhandlers.InsertOrderStatus)
	app.Post("/InsertOrderByRetailer", handlers.CreateRetailerOrderHandler)
	app.Post("/paymentModes", Masterhandlers.InsertPaymentMode)
	app.Post("/InsertLanguages", Masterhandlers.InsertLanguage)
	app.Post("/businessTypeDetails", Masterhandlers.InsertBusinessType)
	app.Post("/listPaymentsDetails", Masterhandlers.InsertPaymentType)
	app.Post("/InsertBusinessCategory", Masterhandlers.InsertBusinessCategory)
	app.Post("/business-branches", Masterhandlers.InsertBusinessBranch)
	app.Post("/InsertDailyPrice", handlers.InsertPriceHandler)
	app.Post("/InsertInvoiceDetails", handlers.InsertInvoiceDetails)
	app.Post("/PostProductCategoryRegional", Masterhandlers.InsertProductCategoryRegional)

	
	// PUT Routes (Updating)
	app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
	app.Put("/categoryUpdate", Masterhandlers.UpdateCategory)
	app.Put("/driverUpdate", Masterhandlers.UpdateDriver)
	app.Put("/locationUpdate", Masterhandlers.UpdateLocation)
	app.Put("/mandiUpdate", Masterhandlers.UpdateMandiDetails)
	app.Put("/languagesUpdate", Masterhandlers.UpdateLanguage)
	app.Put("/productUpdate", Masterhandlers.UpdateProductHandler)
	app.Put("/statesUpdate", Masterhandlers.UpdateState)
	app.Put("/vehicleUpdate", Masterhandlers.UpdateVehicle)
	app.Put("/violationUpdate", Masterhandlers.UpdateMasterViolation)
	app.Put("/usertableUpdate", Masterhandlers.UpdateUser)
	app.Put("/businessUpdate", Masterhandlers.UpdateBusiness)
	app.Put("/orderStatusUpdate", Masterhandlers.UpdateOrderStatus)
	//app.Put("/ordersUpdate", handlers.UpdateOrderHandler)
	app.Put("/dailyPriceUpdate", handlers.UpdatePriceHandler)
	app.Put("/paymentModeUpdate", Masterhandlers.UpdatePaymentMode)
	app.Put("/businessTypeUpdate", Masterhandlers.UpdateBusinessType)
	app.Put("/PaymentTypeUpdate", Masterhandlers.UpdatePaymentType)
	app.Put("/businessCategoryUpdate", Masterhandlers.UpdateBusinessCategory)
	app.Put("/branch/:id", Masterhandlers.UpdateBusinessBranch)
	app.Put("/businessUserDetailesUpdate", Masterhandlers.UpdateBusinessUser) 
	app.Put("/productCategoryRegionalUpdate", Masterhandlers.UpdateProductCategoryRegional) 

}

