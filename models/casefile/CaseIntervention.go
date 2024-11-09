package casefile

import (
	"api/db"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type InterventionFrame string

const (
	TeamDeparture   InterventionFrame = "teamDeparture"
	CaseArrival     InterventionFrame = "caseArrival"
	CaseDeparture   InterventionFrame = "caseDeparture"
	HospitalArrival InterventionFrame = "hospitalArrival"
)

type CaseEmergencyCode int

const (
	RedCode CaseEmergencyCode = iota
	YellowCode
	GreenCode
)

type CaseIntervention struct {
	Reason string `bson:"reason" json:"reason"`

	Code CaseEmergencyCode `bson:"code" json:"code"`

	Address Address `bson:"address" json:"address"`

	CallerPhone string `bson:"callerPhone" json:"callerPhone"`
	CallerName  string `bson:"callername" json:"callerName"`

	TimeFrames struct {
		TeamDeparture   time.Time `bson:"teamDeparture" json:"teamDeparture"`
		CaseArrival     time.Time `bson:"caseArrival" json:"caseArrival"`
		CaseDeparture   time.Time `bson:"caseDeparture" json:"caseDeparture"`
		HospitalArrival time.Time `bson:"hospitalArrival" json:"hospitalArrival"`
	} `bson:"timeFrames" json:"timeFrames"`
}

func (c *Case) SetIntervention(intervention CaseIntervention) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention": intervention,
		},
	})

	if err != nil {
		return
	}

	c.Intervention = intervention

	return
}

func (c *Case) SetInterventionCode(interventionCode CaseEmergencyCode) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.code": interventionCode,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.Code = interventionCode

	return
}

func (c *Case) SetTimeFrame(frame InterventionFrame) (err error) {
	now := time.Now()

	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			fmt.Sprintf("intervention.timeFrames.%v", frame): now,
		},
	})

	if err != nil {
		return
	}

	switch frame {
	case TeamDeparture:
		c.Intervention.TimeFrames.TeamDeparture = now
	case CaseArrival:
		c.Intervention.TimeFrames.CaseArrival = now
	case CaseDeparture:
		c.Intervention.TimeFrames.CaseDeparture = now
	case HospitalArrival:
		c.Intervention.TimeFrames.HospitalArrival = now
	default:
		return errors.New("time frame provided is incorrect")
	}

	return
}

func (c *Case) SetCaller(callerPhone string, callerName string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.callerPhone": callerPhone,
			"intervention.callerName":  callerName,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.CallerPhone = callerPhone
	c.Intervention.CallerName = callerName

	return
}
