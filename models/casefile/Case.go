package casefile

import (
	"api/db"
	"api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Case struct {
	ID           string `bson:"id" json:"id"`
	DispatcherID string `bson:"dispatcherID" json:"dispatcherID"`

	Intervention CaseIntervention `bson:"intervention" json:"intervention"`

	Patient CasePatient `bson:"patient" json:"patient"`

	PrimaryEvaluation   PrimaryEvaluation   `bson:"primaryEvaluation" json:"primaryEvaluation"`
	SecondaryEvaluation SecondaryEvaluation `bson:"secondaryEvaluation" json:"secondaryEvaluation"`

	Procedures Procedures `bson:"procedures" json:"procedures"`

	TeamID string `bson:"teamID" json:"teamID"`

	HospitalID string `bson:"hospitalID" json:"hospitalID"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

func (c *Case) Create(dispatcherID string) (err error) {
	c.ID = utils.GenID(10)
	c.DispatcherID = dispatcherID

	c.Patient = CasePatient{}
	c.Intervention = CaseIntervention{}

	c.CreatedAt = time.Now()

	_, err = db.Cases.InsertOne(db.Ctx, c)

	return
}

func (c *Case) AssignTeam(teamID string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"teamID": teamID,
		},
	})

	if err != nil {
		return
	}

	c.TeamID = teamID

	return
}

func (c *Case) AssignHospital(hospitalID string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"hospitalID": hospitalID,
		},
	})

	if err != nil {
		return
	}

	c.HospitalID = hospitalID

	return
}
