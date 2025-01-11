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
    app.Get("/user-bankdetails", handlers.GetUserBankDetails) 
    log.Fatal(app.Listen(":3000"))
}