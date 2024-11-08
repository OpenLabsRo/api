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

type Dispatcher struct {
	DispatcherID string `bson:"dispatcherID" json:"dispatcherID"`
	AccountID    string `bson:"accountID" json:"accountID"`
}

type DispatcherToken struct {
	Account
	Dispatcher
}

func (d Dispatcher) GenToken(account Account) (token string) {
	claims, _ := sj.ToClaims(DispatcherToken{Account: account, Dispatcher: d})
	claims.SetExpiresAt(time.Now().Add(365 * 24 * time.Hour))

	return claims.Generate(env.JWT_KEY)
}

func (dt *DispatcherToken) ParseToken(token string) (err error) {
	verified := sj.Verify(token, env.JWT_KEY)

	if !verified {
		return errors.New("Could not verify token")
	}

	claims, _ := sj.Parse(token)
	err = claims.Validate()
	claims.ToStruct(&dt)

	return
}

func DispatcherMiddleware(c fiber.Ctx) error {
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

		var dt DispatcherToken
		err := dt.ParseToken(token)
		if err != nil {
			return utils.Error(c, errors.New("an error has occured"))
		}

		c.Locals("id", dt.ID)
		utils.SetLocals(c, "dispatcher", dt)
	}

	if token == "" {
		return utils.Error(c, errors.New("no token provided"))
	}

	return c.Next()
}

func (d *Dispatcher) Get(dispatcherID string) error {
	return db.Dispatchers.FindOne(db.Ctx, bson.M{
		"dispatcherID": dispatcherID,
	}).Decode(&d)
}

func (d *Dispatcher) GetByAccount(accountID string) error {
	return db.Dispatchers.FindOne(db.Ctx, bson.M{
		"accountID": accountID,
	}).Decode(&d)
}
