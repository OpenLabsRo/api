package models

import (
	"api/db"
	"api/env"
	"api/utils"
	"errors"
	"strings"
	"time"

	sj "github.com/brianvoe/sjwt"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Account struct {
	ID string `bson:"id" json:"id"`

	Phone string `bson:"phone" json:"phone"`

	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`

	Passcode string `bson:"passcode" json:"passcode"`
}

func (acc Account) GenToken() (token string) {
	claims, _ := sj.ToClaims(acc)
	claims.SetExpiresAt(time.Now().Add(365 * 24 * time.Hour))

	token = claims.Generate(env.JWT_KEY)
	return
}

func (acc *Account) ParseToken(token string) (err error) {
	verified := sj.Verify(token, env.JWT_KEY)

	if !verified {
		return errors.New("Could not verify token")
	}

	claims, _ := sj.Parse(token)
	err = claims.Validate()
	claims.ToStruct(&acc)

	return
}

func AccountMiddleware(c fiber.Ctx) error {
	var token string

	header := string(c.Get("Authorization"))

	if header != "" &&
		strings.HasPrefix(header, "Bearer") {

		tokens := strings.Fields(header)

		if len(tokens) == 2 {
			token = tokens[1]
		}

		if token == "" {
			return utils.Error(c, errors.New("no token provided"))
		}

		var account Account
		err := account.ParseToken(token)
		if err != nil {
			return utils.Error(c, errors.New("an error has occured"))
		}

		c.Locals("id", account.ID)
		utils.SetLocals(c, "account", account)
	}

	if token == "" {
		return utils.Error(c, errors.New("no token provided"))
	}

	return c.Next()
}

func (acc *Account) Get(id string) error {
	return db.Accounts.FindOne(db.Ctx, bson.M{
		"id": id,
	}).Decode(&acc)

}

func (acc *Account) GetByPhone(phone string) error {
	return db.Accounts.FindOne(db.Ctx, bson.M{
		"phone": phone,
	}).Decode(&acc)
}
