package casefile

import (
	"api/db"
	"api/utils"
)

type Case struct {
	ID           string `bson:"id" json:"id"`
	DispatcherID string `bson:"dispatcherID" json:"dispatcherID"`

	Intervention CaseIntervention `bson:"intervention" json:"intervention"`

	Patient CasePatient `bson:"patient" json:"patient"`
}

func (c *Case) Create(dispatcherID string) (err error) {
	c.ID = utils.GenID(10)
	c.DispatcherID = dispatcherID

	c.Patient = CasePatient{}
	c.Intervention = CaseIntervention{}

	_, err = db.Cases.InsertOne(db.Ctx, c)

	return
}
