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

func StringToFrame(frame string) InterventionFrame {
	switch frame {
	case "teamDeparture":
		return TeamDeparture
	case "caseArrival":
		return CaseArrival
	case "caseDeparture":
		return CaseDeparture
	case "hospitalArrival":
		return HospitalArrival
	default:
		return ""
	}
}

type CaseEmergencyCode int

const (
	RedCode CaseEmergencyCode = iota
	YellowCode
	GreenCode
)

func IntToEmergencyCode(code int) CaseEmergencyCode {
	switch code {
	case 0:
		return RedCode
	case 1:
		return YellowCode
	case 2:
		return GreenCode
	default:
		return -1
	}
}

type CaseSolicitant string

const (
	EmergencyNo CaseSolicitant = "112"
	Family      CaseSolicitant = "family"
	Sanitary    CaseSolicitant = "sanitary"
	Firemen     CaseSolicitant = "firemen"
	Police      CaseSolicitant = "police"
	Militia     CaseSolicitant = "militia"
	Other       CaseSolicitant = "other"
)

type CaseIntervention struct {
	CaseCode           string `bson:"caseCode" json:"caseCode"`
	ActivationProtocol string `bson:"activationProtocol" json:"activationProtocol"`

	Solicitant  CaseSolicitant `bson:"solicitant" json:"solicitant"`
	CallerPhone string         `bson:"callerPhone" json:"callerPhone"`

	EmergencyCode CaseEmergencyCode `bson:"emergencyCode" json:"emergencyCode"`

	Address        Address `bson:"address" json:"address"`
	AddressDetails string  `bson:"addressDetails" json:"addressDetails"`

	Hazards string `bson:"hazards" json:"hazards"`
	Assets  string `bson:"assets" json:"assets"`

	Reason string `bson:"reason" json:"reason"`

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

func (c *Case) SetInterventionSolicitant(solicitant CaseSolicitant) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.solicitant": solicitant,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.Solicitant = solicitant

	return
}

func (c *Case) SetInterventionReason(reason string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.reason": reason,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.Reason = reason
	return
}

func (c *Case) SetInterventionAddress(address Address, addressDetails string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.address":        address,
			"intervention.addressDetails": addressDetails,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.Address = address
	c.Intervention.AddressDetails = addressDetails
	return
}

func (c *Case) SetInterventionCaseCode(caseCode string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.caseCode": caseCode,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.CaseCode = caseCode

	return
}

func (c *Case) SetInterventionActivationProtocol(protocol string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.activationProtocol": protocol,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.ActivationProtocol = protocol

	return
}

func (c *Case) SetInterventionEmergencyCode(code CaseEmergencyCode) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.emergencyCode": code,
		},
	})

	if err != nil {
		return
	}

	switch code {
	case 0:
		c.Intervention.EmergencyCode = RedCode
	case 1:
		c.Intervention.EmergencyCode = YellowCode
	case 2:
		c.Intervention.EmergencyCode = GreenCode
	default:
		c.Intervention.EmergencyCode = -1
	}

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
		},
	})

	if err != nil {
		return
	}

	c.Intervention.CallerPhone = callerPhone

	return
}

func (c *Case) SetInterventionHazards(hazards string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.hazards": hazards,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.Hazards = hazards

	return
}

func (c *Case) SetInterventionAssets(assets string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"intervention.assets": assets,
		},
	})

	if err != nil {
		return
	}

	c.Intervention.Assets = assets

	return
}
