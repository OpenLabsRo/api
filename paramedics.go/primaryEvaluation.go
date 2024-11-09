package paramedics

import (
	"api/models"
	"api/models/casefile"
	"api/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func PrimaryEvaluation(para fiber.Router) {
	primary := para.Group("/primary")

	primary.Post("/mental", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.MentalStatus
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetPrimaryMental(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)

	}, models.ParamedicMiddleware)

	primary.Post("/airways", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.Airways
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetPrimaryAirways(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)

	}, models.ParamedicMiddleware)

	primary.Post("/breathing", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.Breathing
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetPrimaryBreathing(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)

	}, models.ParamedicMiddleware)

	primary.Post("/circulation", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.Circulation
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetPrimaryCirculation(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)

	}, models.ParamedicMiddleware)
}