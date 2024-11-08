package main

import (
	"api/db"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	db.InitDB()

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("PONG\n")
	})

	app.Listen(":6666")
}
