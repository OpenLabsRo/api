package accounts

import "github.com/gofiber/fiber/v3"

func Endpoints(app *fiber.App) {
	acc := app.Group("/accounts")

	acc.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("PONG\n")
	})
}
