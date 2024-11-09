package models

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Ambulance struct {
	ID     string `bson:"id" json:"id"`
	TeamID string `bson:"teamID" json:"teamID"`

	Available bool    `bson:"available" json:"available"`
	Lat       float64 `bson:"lat" json:"lat"`
	Lng       float64 `bson:"lng" json:"lng"`
}

func UpdateAmbulancePost(ambulanceID string, lat float64, lng float64) error {
	_, err := db.Ambulances.UpdateOne(db.Ctx, bson.M{
		"id": ambulanceID,
	}, bson.M{
		"$set": bson.M{
			"lat": lat,
			"lng": lng,
		},
	})

	return err
}

func GetAmbulances() (ambulances []Ambulance, err error) {
	cursor, err := db.Ambulances.Find(db.Ctx, bson.M{})
	if err != nil {
		return
	}

	err = cursor.All(db.Ctx, &ambulances)
	return
}
