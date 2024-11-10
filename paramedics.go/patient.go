package paramedics

import (
	"api/models/casefile"
	"api/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func patient(para fiber.Router) {
	patient := para.Group("/patient")

	patient.Patch("/cnp", func(c fiber.Ctx) error {
		id := c.Query("id")
		cnp := c.Query("cnp")

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetPatientCNP(cnp)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(bson.M{
			"cnp": cnp,
		})
	})
	patient.Patch("/name", func(c fiber.Ctx) error {
		id := c.Query("id")
		firstName := c.Query("firstName")
		lastName := c.Query("lastName")

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetPatientName(firstName, lastName)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(bson.M{
			"firstName": firstName,
			"lastName":  lastName,
		})
	})

	patient.Patch("/age", func(c fiber.Ctx) error {
		id := c.Query("id")
		age, _ := strconv.Atoi(c.Query("age"))
		birthdate := c.Query("birthdate")

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetPatientAge(birthdate, age)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(bson.M{
			"birthdate": birthdate,
			"age":       age,
		})
	})

	patient.Patch("/weight", func(c fiber.Ctx) error {
		id := c.Query("id")
		weight, _ := strconv.Atoi(c.Query("weight"))

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetPatientWeight(weight)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(bson.M{
			"weight": weight,
		})
	})

	patient.Patch("/sex", func(c fiber.Ctx) error {
		id := c.Query("id")
		sex := c.Query("sex")

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetPatientSex(sex)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(bson.M{
			"sex": sex,
		})
	})
	patient.Patch("/home-address", func(c fiber.Ctx) error {
		id := c.Query("id")
		homeAddress := c.Query("homeAddress")

		var caseFile casefile.Case
		caseFile.ID = id

		err := caseFile.SetPatientHomeAddress(homeAddress)
		if err != nil {
			return utils.Error(c, err)
		}

		return c.JSON(bson.M{
			"homeAddress": homeAddress,
		})
	})
}
