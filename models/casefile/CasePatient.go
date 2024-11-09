package casefile

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type CasePatient struct {
	CNP string `bson:"cnp" json:"cnp"`

	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`

	Age         int    `bson:"age" json:"age"`
	DateOfBirth string `bson:"dateOfBirth" json:"dateOfBirth"`

	Weight int    `bson:"weight" json:"weight"`
	Sex    string `bson:"sex" json:"sex"`

	HomeAddress Address `bson:"homeAddress" json:"homeAddress"`

	Allergies string `bson:"allergies" json:"allergies"`
	Diseases  string `bson:"diseases" json:"diseases"`
}

func (c *Case) SetPatient(patient CasePatient) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"patient": patient,
		},
	})

	if err != nil {
		return
	}

	c.Patient = patient

	return
}
