package accounts

import (
	"api/models"
	"api/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Endpoints(app *fiber.App) {
	acc := app.Group("/accounts")

	acc.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("PONG\n")
	})

	acc.Get("/whoami", func(c fiber.Ctx) error {
		var account models.Account
		utils.GetLocals(c, "account", &account)

		return c.JSON(account)
	}, models.AccountMiddleware)

	acc.Post("/login", func(c fiber.Ctx) error {
		var body struct {
			Phone    string `json:"phone"`
			Passcode string `json:"passcode"`
		}
		json.Unmarshal(c.Body(), &body)

		var account models.Account
		err := account.GetByPhone(body.Phone)
		if err != nil {
			return utils.Error(c, err)
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(account.Passcode),
			[]byte(body.Passcode),
		)
		if err != nil {
			return utils.Error(c, err)
		}

		token := account.GenToken()

		return c.JSON(
			bson.M{
				"token":   token,
				"account": account,
			},
		)
	})
}
