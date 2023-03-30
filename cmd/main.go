package main

import (
    "github.com/gofiber/fiber/v2"
    "https://github.com/dr942120/RestAPI_Go"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    setupRoutes(app)

    app.Listen(":3000")
}