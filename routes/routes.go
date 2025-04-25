package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/handlers"
	Masterhandlers "github.com/PragaL15/go_newBackend/handlers/master"
	TrendHandlers "github.com/PragaL15/go_newBackend/handlers/trendHandlers"
	Marketoppurtinities "github.com/PragaL15/go_newBackend/handlers/marketOppurtinities"
	RestockingStock "github.com/PragaL15/go_newBackend/handlers/restocking"
)

func RegisterRoutes(app *fiber.App) {
	// Registering routes for GET methods
	app.Get("/getUsers", Masterhandlers.GetAllUsers)
	app.Get("/getCategoriesBySupID/:super_cat_id", Masterhandlers.GetCategoriesBySuperCatID)
	app.Get("/getDrivers", Masterhandlers.GetDrivers)
	app.Get("/getLocations", Masterhandlers.GetLocations)
	app.Get("/getMandiDetails", Masterhandlers.GetAllMandiDetails)
	app.Get("/getAllLanguages", Masterhandlers.GetAllLanguages)
	app.Get("/getProducts", Masterhandlers.GetAllProducts)
	app.Get("/getStates", Masterhandlers.GetAllStates)
	app.Get("/getBusinesses", Masterhandlers.GetAllBusinesses)
	app.Get("/getOrderStatus", Masterhandlers.GetAllOrderStatuses)
	app.Get("/getBusinessTypes", Masterhandlers.GetBusinessTypes)
	app.Get("/getModeOfPayments", Masterhandlers.GetPaymentModes)
	app.Get("/getListPaymentMethod", Masterhandlers.GetAllCashPaymentsType)
	app.Get("/getOrderHistoryDetails", handlers.GetOrderHistory)
	app.Get("/getDailyPriceDetails", handlers.GetPriceHandler)
	app.Get("/getCashPaymentTypes", Masterhandlers.GetAllCashPaymentsType)
	app.Get("/getBusinessCategory", Masterhandlers.GetBusinessCategories)
	app.Get("/getBusinessBranches", Masterhandlers.GetAllBusinessBranches)
	app.Get("/getAllOrderDetails/:id", handlers.GetOrderDetailsHandler)
	app.Get("/getCartitems/:id", handlers.GetCart)
	app.Get("/getBusinessUsers", Masterhandlers.GetAllBusinessUsers)
	app.Get("/getCategoryRegionalName", Masterhandlers.GetProductCategoryRegional)
	app.Get("/getProductRegionalName", Masterhandlers.GetAllProductRegionalNames)

	app.Get("/getSalesValue/monthly", TrendHandlers.GetSalesMonthlyHandler)  
	app.Get("/getSalesValue/weekly", TrendHandlers.GetSalesWeeklyHandler)    
	app.Get("/getSalesValue/yearly", TrendHandlers.GetSalesYearlyHandler) 
	//app.Get("/getTopSellingDaily", TrendHandlers.GetTopSellingDailyHandler) 
	app.Get("/getTopSellingWeekly", TrendHandlers.GetTopSellingWeeklyHandler) 
	app.Get("/getTopSellingMonthly", TrendHandlers.GetTopSellingMonthlyHandler) 
	app.Get("/getTopSellingYearly", TrendHandlers.GetTopSellingYearlyHandler) 
	app.Get("/getMandiStockedProduct", TrendHandlers.GetLeastStockedProductsHandler) 
	app.Get("/getLowStockItems", TrendHandlers.GetLowStockItemsHandler) 
	app.Get("/getStockAvailability", TrendHandlers.GetStockAvailabilityPercentageHandler) 
	app.Get("/getCurrentStockByMandi/:mandi_id", TrendHandlers.GetCurrentStockByMandiHandler)
	app.Get("/getAllBulkOrderDetails", Marketoppurtinities.GetAllBulkOrderDetailsHandler)
	app.Get("/getTopRetailerDetails", Marketoppurtinities.GetTopRetailerHandler)
	app.Get("/getReStockProductsHandler", RestockingStock.GetLowStockProductsHandler)

	app.Get("/getWholesellerPriceComparison", TrendHandlers.GetWholesellerPriceComparisonHandler)
//http://localhost:3000/getWholesellerPriceComparison?product_ids=2,3


	// Registering routes for GET methods with ID
	app.Get("/getProducts/:product_id", Masterhandlers.GetProductByID)
	app.Get("/getUsers/:id", Masterhandlers.GetUserByID)
	app.Get("/getBusinesses/:id", Masterhandlers.GetBusinessByID)
	app.Get("/getMandiDetails/:id", Masterhandlers.GetMandiDetailsByID)
	app.Get("/getCategories/:category_id", Masterhandlers.GetCategoryByID)
	app.Get("/getStates/:state_id", Masterhandlers.GetStateByID)
	app.Get("/getProductByCatId/:category_id", Masterhandlers.GetProductsByCategoryID)
	app.Get("/getPaymentModeById/:id", Masterhandlers.GetPaymentModeByID)
	app.Get("/getOrderStatusById/:id", Masterhandlers.GetOrderStatusByID)
	app.Get("/getBusinessTypeById/:id", Masterhandlers.GetBusinessTypeByID)
	app.Get("/getDriverById/:id", Masterhandlers.GetDriverByID)
	app.Get("/getBusinessesbranch/:id", Masterhandlers.GetBusinessBranchByID)
	app.Get("/getBusinessUser/:id", Masterhandlers.GetBusinessUserByID) 
	app.Get("/getProductCategoryRegional/:id", Masterhandlers.GetProductCategoryRegionalByID) 
	app.Get("/invoice/:invoice_id", handlers.GetInvoiceDetails)

	// POST Routes
	app.Post("/categoryDetails", Masterhandlers.InsertCategory)
	app.Post("/locationDetails", Masterhandlers.InsertLocation)
	app.Post("/mandiDetails", Masterhandlers.InsertMandiDetails)
	app.Post("/productDetails", Masterhandlers.InsertProduct)
	app.Post("/stateDetails", Masterhandlers.InsertState)
	app.Post("/userTableDetails", Masterhandlers.InsertUser)
	app.Post("/businessDetails", Masterhandlers.InsertBusiness)
	app.Post("/orderStatusDetails", Masterhandlers.InsertOrderStatus)
	app.Post("/InsertOrderByRetailer", handlers.CreateRetailerOrderHandler)
	app.Post("/paymentModes", Masterhandlers.InsertPaymentMode)
	app.Post("/InsertLanguages", Masterhandlers.InsertLanguage)
	app.Post("/businessTypeDetails", Masterhandlers.InsertBusinessType)
	app.Post("/listCashPaymentsDetails", Masterhandlers.InsertCashPaymentType)
	app.Post("/InsertBusinessCategory", Masterhandlers.InsertBusinessCategory)
	app.Post("/business-branches", Masterhandlers.InsertBusinessBranch)
	app.Post("/InsertDailyPrice", handlers.InsertPriceHandler)
	app.Post("/InsertInvoiceDetails", handlers.InsertInvoiceDetails)
	app.Post("/InsertCartDetails", handlers.InsertCart)
	app.Post("/PostCategoryRegional", Masterhandlers.InsertProductCategoryRegional)
	app.Post("/PostProductRegional", Masterhandlers.InsertProductRegionalName)

	// PUT Routes
	app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
	app.Put("/categoryUpdate", Masterhandlers.UpdateCategory)
	app.Put("/driverUpdate", Masterhandlers.UpdateDriver)
	app.Put("/locationUpdate", Masterhandlers.UpdateLocation)
	app.Put("/mandiUpdate", Masterhandlers.UpdateMandiDetails)
	app.Put("/languagesUpdate", Masterhandlers.UpdateLanguage)
	app.Put("/productUpdate", Masterhandlers.UpdateProductHandler)
	app.Put("/statesUpdate/:id", Masterhandlers.UpdateState)
	app.Put("/vehicleUpdate", Masterhandlers.UpdateVehicle)
	app.Put("/violationUpdate", Masterhandlers.UpdateMasterViolation)
	app.Put("/userUpdate/:id", Masterhandlers.UpdateUser)
	app.Put("/businessUpdate", Masterhandlers.UpdateBusiness)
	app.Put("/orderStatusUpdate/:id", Masterhandlers.UpdateOrderStatus)
	app.Put("/dailyPriceUpdate", handlers.UpdatePriceHandler)
	app.Put("/paymentModeUpdate", Masterhandlers.UpdatePaymentMode)
	app.Put("/businessTypeUpdate", Masterhandlers.UpdateBusinessType)
	app.Put("/CashPaymentTypeUpdate", Masterhandlers.UpdateCashPaymentType)
	app.Put("/businessCategoryUpdate", Masterhandlers.UpdateBusinessCategory)
	app.Put("/branchDetailsUpdate", Masterhandlers.UpdateBusinessBranch)
	app.Put("/businessUserDetailesUpdate", Masterhandlers.UpdateBusinessUser) 
	app.Put("/CategoryRegionalUpdate", Masterhandlers.UpdateProductCategoryRegional) 
	app.Put("/ProductRegionalUpdate", Masterhandlers.UpdateProductRegionalName) 

	// DELETE Routes
	app.Delete("/:cart_id/items/:product_id", handlers.DeleteCartItem) 
}
