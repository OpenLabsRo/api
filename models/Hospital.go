package models

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Hospital struct {
	ID        string `bson:"id" json:"id"`
	Name      string `bson:"name" json:"name"`
	Available bool   `bson:"available" json:"available"`
}

func GetHospital() (hospitals []Hospital, err error) {
	cursor, err := db.Hospitals.Find(db.Ctx, bson.M{})
	if err != nil {
		return
	}

	if err = cursor.All(db.Ctx, &hospitals); err != nil {
		return
	}

	return
}
