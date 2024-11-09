package dispatchers

import "github.com/gofiber/fiber/v3"

func Endpoints(app *fiber.App) {
	dis := app.Group("/dispatchers")

	dis.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("PONG\n")
	})

}
