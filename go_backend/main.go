package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "go-fiber-stored-procedure/db"
    "go-fiber-stored-procedure/handlers"
)

func main() {
    // Connect to the database
    db.ConnectDB()

    // Create a new Fiber app
    app := fiber.New()

    // Define routes
    app.Put("/update-salary", handlers.UpdateEmployeeSalary)
    app.Get("/users", handlers.GetAllUsers) // New route to fetch users

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
