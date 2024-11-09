package dispatchers

import (
	"api/models"
	"api/utils"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func Endpoints(app *fiber.App) {
	dis := app.Group("/dispatchers")

	dis.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("PONG\n")
	})

	dis.Get("/whoami", func(c fiber.Ctx) error {
		var dis models.Dispatcher
		utils.GetLocals(c, "dispatcher", &dis)

		return c.JSON(dis)
	})

	dis.Post("/auth", func(c fiber.Ctx) error {
		var account models.Account
		utils.GetLocals(c, "account", &account)

		var dispatcher models.Dispatcher
		err := dispatcher.GetByAccount(account.ID)
		if err != nil {
			return utils.Error(c, err)
		}

		token := dispatcher.GenToken(account)
		return c.JSON(bson.M{
			"dispatcher": dispatcher,
			"account":    account,
			"token":      token,
		})
	}, models.AccountMiddleware)

}
