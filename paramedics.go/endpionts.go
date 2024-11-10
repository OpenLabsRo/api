package paramedics

import (
	"api/models"
	"api/models/casefile"
	"api/utils"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func Endpoints(app *fiber.App) {
	para := app.Group("/paramedics")

	para.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("PONG\n")
	})

	para.Get("/whoami", func(c fiber.Ctx) error {
		var paramedic models.Paramedic
		utils.GetLocals(c, "paramedic", &paramedic)

		return c.JSON(paramedic)
	}, models.ParamedicMiddleware)

	para.Post("/auth", func(c fiber.Ctx) error {
		var account models.Account
		utils.GetLocals(c, "account", &account)

		var paramedic models.Paramedic
		err := paramedic.GetByAccount(account.ID)
		if err != nil {
			return utils.Error(c, err)
		}

		token := paramedic.GenToken(account)
		return c.JSON(bson.M{
			"paramedic": paramedic,
			"account":   account,
			"token":     token,
		})
	}, models.AccountMiddleware)

	para.Patch("/case/intervention/timeframes", func(c fiber.Ctx) error {
		id := c.Query("id")
		frame := c.Query("frame")

		var body struct {
			Time time.Time `json:"time"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetTimeFrame(casefile.StringToFrame(frame))
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

	primaryEvaluation(para)
	secondaryEvaluation(para)
	procedures(para)
}
