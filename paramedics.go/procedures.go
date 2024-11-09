package paramedics

import (
	"api/models"
	"api/models/casefile"
	"api/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func procedures(para fiber.Router) {
	procedures := para.Group("/procedures")

	procedures.Patch("/intubation", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.ProcedureIntubation
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetProceduresIntubation(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

	procedures.Patch("/ventilation", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.ProcedureVentilation
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetProceduresVentilation(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

	procedures.Patch("/intravenous", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.ProcedureIntravenous
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetProceduresIntravenous(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

	procedures.Patch("/intraosseous", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.ProcedureIntraosseous
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetProceduresIntraosseous(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

	procedures.Patch("/trauma", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.ProcedureTrauma
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetProceduresTrauma(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

	procedures.Patch("/hemostasis", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.ProcedureHemostasis
		json.Unmarshal(c.Body(), &body)

		err := caseFile.SetProceduresHemostasis(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

	procedures.Patch("/medication", func(c fiber.Ctx) error {
		id := c.Query("id")

		var caseFile casefile.Case
		caseFile.ID = id

		var body casefile.ProcedureMedication
		json.Unmarshal(c.Body(), &body)

		err := caseFile.AddProceduresMedication(body)

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.ParamedicMiddleware)

}
