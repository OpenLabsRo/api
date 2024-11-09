package casefile

import (
	"api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type CasePatient struct {
	CNP string `bson:"cnp" json:"cnp"`

	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`

	Age       int    `bson:"age" json:"age"`
	Birthdate string `bson:"birthdate" json:"birthdate"`

	Weight int    `bson:"weight" json:"weight"`
	Sex    string `bson:"sex" json:"sex"`

	HomeAddress string `bson:"homeAddress" json:"homeAddress"`

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

func (c *Case) SetPatientCNP(cnp string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"patient.cnp": cnp,
		},
	})

	if err != nil {
		return
	}

	c.Patient.CNP = cnp

	return
}

func (c *Case) SetPatientName(firstName string, lastName string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"patient.lastName":  firstName,
			"patient.firstName": lastName,
		},
	})

	if err != nil {
		return
	}

	c.Patient.FirstName = firstName
	c.Patient.LastName = lastName

	return
}

func (c *Case) SetPatientAge(birthdate string, age int) (err error) {
	if birthdate != "" {
		_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
			"id": c.ID,
		}, bson.M{
			"$set": bson.M{
				"patient.birthdate": birthdate,
			},
		})

		if err != nil {
			return
		}

		c.Patient.Birthdate = birthdate
	} else {
		_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
			"id": c.ID,
		}, bson.M{
			"$set": bson.M{
				"patient.birthdate": birthdate,
			},
		})

		if err != nil {
			return
		}

		c.Patient.Age = age
	}

	return
}

func (c *Case) SetPatientWeight(weight int) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"patient.weight": weight,
		},
	})

	if err != nil {
		return
	}

	c.Patient.Weight = weight
	return
}

func (c *Case) SetPatientSex(sex string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"patient.sex": sex,
		},
	})

	if err != nil {
		return
	}

	c.Patient.Sex = sex

	return
}

func (c *Case) SetPatientHomeAddress(homeAddress string) (err error) {
	_, err = db.Cases.UpdateOne(db.Ctx, bson.M{
		"id": c.ID,
	}, bson.M{
		"$set": bson.M{
			"patient.homeAddress": homeAddress,
		},
	})

	if err != nil {
		return
	}

	c.Patient.HomeAddress = homeAddress

	return
}
