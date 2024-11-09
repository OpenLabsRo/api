package casefile

import (
	"api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ProcedureIntubation struct {
	Yes struct {
		OrotrachealTube bool `bson:"orotrachealTube" json:"orotrachealTube"`
		NasalCannula    bool `bson:"nasalCannula" json:"nasalCannula"`
		Supraglottic    bool `bson:"supraglottic" json:"supraglottic"`
	} `bson:"yes" json:"yes"`
	No bool `bson:"no" json:"no"`
}

type ProcedureVentilation struct {
	Mech   bool `bson:"mech" json:"mech"`
	Baloon bool `bson:"baloon" json:"baloon"`
}

type ProcedureIntravenous struct {
	Central   bool `bson:"central" json:"central"`
	Periferic bool `bson:"periferic" json:"periferic"`
}

type ProcedureIntraosseous struct {
	Femoral bool `bson:"femoral" json:"femoral"`
	Tibial  bool `bson:"tibial" json:"tibial"`
}

type ProcedureTrauma struct {
	Extricated     bool `bson:"extricated" json:"extricated"`
	CervicalCollar bool `bson:"cervicalCollar" json:"cervicalCollar"`
	VacuumSplints  bool `bson:"vacuumSplints" json:"vacuumSplints"`
	Kendrick       bool `bson:"kendrick" json:"kendrick"`
	Shovel         bool `bson:"shovel" json:"shovel"`
}

type ProcedureHemostasis struct {
	Gauze      bool   `bosn:"gauze" json:"gauze"`
	Compress   bool   `bosn:"compress" json:"compress"`
	Tourniquet bool   `bosn:"tourniquet" json:"tourniquet"`
	Other      string `bson:"other" json:"other"`
}

type ProcedureMedication struct {
	Name           string    `bson:"name" json:"name"`
	Dosage         string    `bson:"dosage" json:"dosage"`
	AdministeredAt time.Time `bson:"administeredAt" json:"administeredAt"`
}

type Procedures struct {
	Intubation    ProcedureIntubation   `bson:"intubation" json:"intubation"`
	Ventilation   ProcedureVentilation  `bson:"ventilation" json:"ventilation"`
	Intravenous   ProcedureIntravenous  `bson:"intravenous" json:"intravenous"`
	Intraosseuous ProcedureIntraosseous `bson:"intraosseous" json:"intraosseus"`
	Trauma        ProcedureTrauma       `bson:"trauma" json:"trauma"`
	Hemostasis    ProcedureHemostasis   `bson:"hemostasis" json:"hemostasis"`
	Medications   []ProcedureMedication `bson:"medications" json:"medications"`
}

func (c *Case) SetProceduresIntubation(intubation ProcedureIntubation) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"procedures.intubation": intubation,
		},
	})

	if err != nil {
		return
	}

	c.Procedures.Intubation = intubation

	return
}

func (c *Case) SetProceduresVentilation(ventilation ProcedureVentilation) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"procedures.ventilation": ventilation,
		},
	})

	if err != nil {
		return
	}

	c.Procedures.Ventilation = ventilation

	return
}

func (c *Case) SetProceduresIntravenous(intravenous ProcedureIntravenous) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"procedures.intravenous": intravenous,
		},
	})

	if err != nil {
		return
	}

	c.Procedures.Intravenous = intravenous

	return
}

func (c *Case) SetProceduresIntraosseous(intraosseous ProcedureIntraosseous) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"procedures.intraosseous": intraosseous,
		},
	})

	if err != nil {
		return
	}

	c.Procedures.Intraosseuous = intraosseous

	return
}

func (c *Case) SetProceduresTrauma(trauma ProcedureTrauma) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"procedures.trauma": trauma,
		},
	})

	if err != nil {
		return
	}

	c.Procedures.Trauma = trauma

	return
}

func (c *Case) SetProceduresHemostasis(hemostasis ProcedureHemostasis) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"procedures.hemostasis": hemostasis,
		},
	})

	if err != nil {
		return
	}

	c.Procedures.Hemostasis = hemostasis

	return
}

func (c *Case) AddProceduresMedication(medication ProcedureMedication) (err error) {
	medication.AdministeredAt = time.Now()

	err = db.Cases.FindOne(db.Ctx, bson.M{
		"id": c.ID,
	}).Decode(&c)

	if err != nil {
		return
	}

	newMedications := append(c.Procedures.Medications, medication)

	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"procedures.medications": newMedications,
		},
	})

	if err != nil {
		return
	}

	c.Procedures.Medications = newMedications

	return
}
