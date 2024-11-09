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

type Paramedic struct {
	ParamedicID string `bson:"paramedicID" json:"paramedicID"`
	AccountID   string `bson:"accountID" json:"accountID"`

	Certified bool `bson:"certified" json:"certified"`

	TeamID string `bson:"teamID" json:"teamID"`
}

type ParamedicToken struct {
	Account
	Paramedic
}

func (p Paramedic) GenToken(account Account) (token string) {
	claims, _ := sj.ToClaims(ParamedicToken{Account: account, Paramedic: p})
	claims.SetExpiresAt(time.Now().Add(365 * 24 * time.Hour))

	return claims.Generate(env.JWT_KEY)
}

func (pt *ParamedicToken) ParseToken(token string) (err error) {
	verified := sj.Verify(token, env.JWT_KEY)

	if !verified {
		return errors.New("Could not verify token")
	}

	claims, _ := sj.Parse(token)
	err = claims.Validate()
	claims.ToStruct(&pt)

	return
}

func ParamedicMiddleware(c fiber.Ctx) error {
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

		var pt ParamedicToken
		err := pt.ParseToken(token)
		if err != nil {
			return utils.Error(c, errors.New("an error has occured"))
		}

		c.Locals("id", pt.ID)
		utils.SetLocals(c, "paramedic", pt)
	}

	if token == "" {
		return utils.Error(c, errors.New("no token provided"))
	}

	return c.Next()
}

func (p *Paramedic) Get(paramedicID string) error {
	return db.Dispatchers.FindOne(db.Ctx, bson.M{
		"paramedicID": paramedicID,
	}).Decode(&p)
}

func (p *Paramedic) GetByAccount(accountID string) error {
	return db.Dispatchers.FindOne(db.Ctx, bson.M{
		"accountID": accountID,
	}).Decode(&p)
}

func GetParamedics() (paramedics []Paramedic, err error) {
	cursor, err := db.Paramedics.Find(db.Ctx, bson.M{})
	if err != nil {
		return
	}

	err = cursor.All(db.Ctx, &paramedics)
	return
}
