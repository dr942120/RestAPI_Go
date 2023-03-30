// cmd/routes.go

package main

import (
    "github.com/gofiber/fiber/v2"
    "https://github.com/dr942120/RestAPI_Go"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", handlers.Home)
}
