package paramedics

import (
	"api/models"
	"api/models/casefile"
	"api/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func secondaryEvaluation(para fiber.Router) {
	secondary := para.Group("/secondary")

	secondary.Patch("/usual-symptoms", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.UsualSymptoms
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetSecondaryUsual(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

	secondary.Patch("/trauma", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.Trauma
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetSecondaryTrauma(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)
}
