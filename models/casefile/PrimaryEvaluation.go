package casefile

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type MentalStatus struct {
	Conscious struct {
		Normal  bool `bson:"normal" json:"normal"`
		Altered bool `bson:"altered" json:"altered"`
	} `bson:"conscious" json:"conscious"`
	Unconscious bool `bson:"unconscious" json:"unconscious"`
}

type Airways struct {
	Open       bool `bson:"open" json:"open"`
	Obstructed struct {
		Fully     bool `bson:"fully" json:"fully"`
		Partially bool `bson:"partially" json:"partially"`
	} `bson:"obstructed" json:"obstructed"`
}

type Breathing struct {
	Normal    bool `bson:"normal" json:"normal"`
	Absent    bool `bson:"absent" json:"absent"`
	Dispnea   bool `bson:"dispnea" json:"dispnea"`
	Intubated struct {
		GuedelPipe bool `bson:"guedelPipe" json:"guedelPipe"`
		Tube       bool `bson:"tube" json:"tube"`
		Ventilator bool `bson:"ventilator" json:"ventilator"`
		Baloon     bool `bson:"baloon" json:"baloon"`
	} `bson:"intubated" json:"intubated"`
}

type Circulation struct {
	Present struct {
		Full       bool `bson:"full" json:"full"`
		Filiform   bool `bson:"filiform" json:"filiform"`
		Rhythmic   bool `bson:"rhythmic" json:"rhythmic"`
		Arrhythmic bool `bson:"arrhythmic" json:"arrhythmic"`
	} `bson:"present" json:"present"`
	Absent bool `bson:"absent" json:"absent"`
}

type PrimaryEvaluation struct {
	MentalStatus MentalStatus `bson:"mentalStatus" json:"mentalStatus"`
	Airways      Airways      `bson:"airways" json:"airways"`
	Breathing    Breathing    `bson:"breathing" json:"breathing"`
	Circulation  Circulation  `bson:"circulation" json:"circulation"`
}

func (c *Case) SetPrimaryMental(mentalStatus MentalStatus) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"primaryEvaluation.mentalStatus": mentalStatus,
		},
	})

	if err != nil {
		return
	}

	c.PrimaryEvaluation.MentalStatus = mentalStatus

	return
}

func (c *Case) SetPrimaryAirways(airways Airways) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"primaryEvaluation.airways": airways,
		},
	})

	if err != nil {
		return
	}

	c.PrimaryEvaluation.Airways = airways

	return
}

func (c *Case) SetPrimaryBreathing(breathing Breathing) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"primaryEvaluation.breathing": breathing,
		},
	})

	if err != nil {
		return
	}

	c.PrimaryEvaluation.Breathing = breathing

	return
}

func (c *Case) SetPrimaryCirculation(circulation Circulation) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"primaryEvaluation.circulation": circulation,
		},
	})

	if err != nil {
		return
	}

	c.PrimaryEvaluation.Circulation = circulation

	return
}
