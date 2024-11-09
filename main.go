package main

import (
	"api/accounts"
	"api/dispatchers"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// db.InitDB()

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("PONG\n")
	})

	accounts.Endpoints(app)
	dispatchers.Endpoints(app)

	app.Listen(":6666")
}
