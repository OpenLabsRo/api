package dispatchers

import (
	"api/models"
	"api/models/casefile"
	"api/utils"
	"encoding/json"
	"fmt"

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

	dis.Post("/case", func(c fiber.Ctx) error {
		var caseFile casefile.Case
		err := caseFile.Create(fmt.Sprintf("%v", c.Locals("dispatcherID")))

		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(caseFile)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/reason", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			Reason string `json:"reason"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetInterventionReason(body.Reason)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/address", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			Address        string `json:"address"`
			AddressDetails string `json:"addressDetails"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetInterventionAddress(body.Address, body.AddressDetails)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/emergency-code", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			Code casefile.CaseEmergencyCode `json:"code"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetInterventionEmergencyCode(body.Code)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/case-code", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			CaseCode string `json:"caseCode"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetInterventionCaseCode(body.CaseCode)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/activation-protocol", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			ActivationProtocol string `json:"activationProtocol"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetInterventionActivationProtocol(body.ActivationProtocol)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/caller", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			CallerPhone string `json:"callerPhone"`
			CallerName  string `json:"callerName"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetCaller(body.CallerPhone, body.CallerName)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/solicitant", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			Solicitant casefile.CaseSolicitant `json:"solicitant"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetInterventionSolicitant(body.Solicitant)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/assets", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			Assets string `json:"assets"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetInterventionAssets(body.Assets)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)

	dis.Patch("/case/intervention/hazards", func(c fiber.Ctx) error {
		id := c.Query("id")

		var body struct {
			Hazards string `json:"hazards"`
		}
		json.Unmarshal(c.Body(), &body)

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetInterventionAssets(body.Hazards)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(body)
	}, models.DispatcherMiddleware)
}
