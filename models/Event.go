package models

import (
	"api/db"
	"api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type EventType string

const (
	RequestSupport  EventType = "support"
	RequestHospital EventType = "hospital"
	RequestPolice   EventType = "Police"
)

type Event struct {
	ID        string    `bson:"id" json:"id"`
	CaseID    string    `bson:"caseID" json:"caseID"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	Type      string    `bson:"type" json:"type"`
	Handled   bool      `bson:"handled" json:"handled"`
}

func (e *Event) Init(caseID string) {
	e.ID = utils.GenID(15)
	e.CaseID = caseID
	e.CreatedAt = time.Now()
	e.Handled = false
}

func (e *Event) Create(caseID string, eType EventType) (err error) {
	e.Init(caseID)
	e.Type = string(eType)

	_, err = db.Events.InsertOne(db.Ctx, e)

	return
}

func GetEvents() (events []Event, err error) {
	cursor, err := db.Events.Find(db.Ctx, bson.M{})
	if err != nil {
		return
	}

	err = cursor.All(db.Ctx, &events)

	return
}
