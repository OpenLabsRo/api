package casefile

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type UsualSymptoms struct {
	Nausea       bool   `bson:"nausea" json:"nausea"`
	Vomiting     bool   `bson:"vomiting" json:"vomiting"`
	Perspiration bool   `bson:"perspiration" json:"perspiration"`
	Dizziness    bool   `bson:"dizziness" json:"dizziness"`
	Pain         string `bson:"pain" json:"pain"`
}

type Trauma struct {
	Wound     bool `bson:"wound" json:"wound"`
	Contusion bool `bson:"contusion" json:"contusion"`
	Fracture  struct {
		Open   bool `bson:"open" json:"open"`
		Closed bool `bson:"closed" json:"closed"`
	} `bson:"fracture" json:"fracture"`
	Burn struct {
		AirwaysAffected bool `bson:"airwaysAffected" json:"airwaysAffected"`
		Flame           bool `bson:"flame" json:"flame"`
		Liquids         bool `bson:"liquids" json:"liquids"`
		Solids          bool `bson:"solids" json:"solids"`
		Vapors          bool `bson:"vapors" json:"vapors"`
		Gas             bool `bson:"gas" json:"gas"`
		Chemical        bool `bson:"chemical" json:"chemical"`
	} `bson:"burn" json:"burn"`

	Hypothemia bool `bson:"hypothermia" json:"hypothermia"`
	Drowning   bool `bson:"drowning" json:"drowning"`
}

type SecondaryEvaluation struct {
	UsualSymptoms UsualSymptoms `bson:"usualSymptoms" json:"usualSymptoms"`
	Trauma        Trauma        `bson:"trauma" json:"trauma"`
}

func (c *Case) SetSecondaryUsual(usual UsualSymptoms) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"secondaryEvaluation.usualSymptoms": usual,
		},
	})

	if err != nil {
		return
	}

	c.SecondaryEvaluation.UsualSymptoms = usual

	return
}

func (c *Case) SetSecondaryTrauma(trauma Trauma) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"secondaryEvaluation.trauma": trauma,
		},
	})

	if err != nil {
		return
	}

	c.SecondaryEvaluation.Trauma = trauma

	return
}
