package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/PragaL15/go_newBackend/handlers"
)
func main() {
    db.ConnectDB()
    app := fiber.New()
    app.Get("/users", handlers.GetAllUsers) 
    app.Post("/user-bank-details", handlers.InsertUserBankDetail)
    app.Put("/user-bank-details", handlers.UpdateUserBankDetail)
    app.Delete("/user-bank-details/:id", handlers.DeleteUserBankDetail)
    log.Fatal(app.Listen(":3000"))
}