package models

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Team struct {
	ID string `bson:"team" json:"team"`

	Paramedics []Paramedic `bson:"paramedics" json:"paramedics"`

	AmbulanceID string `bson:"ambulance" json:"ambulanceID"`

	Certified bool `bson:"certified" json:"certified"`
}

func (t *Team) Get(teamID string) error {
	return db.Teams.FindOne(db.Ctx, bson.M{
		"id": teamID,
	}).Decode(&t)
}

func GetTeams() (teams []Team, err error) {
	cursor, err := db.Teams.Find(db.Ctx, bson.M{})
	if err != nil {
		return
	}

	err = cursor.All(db.Ctx, &teams)
	return
}
