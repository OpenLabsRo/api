package casefile

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
	OtherCode
)

type CaseIntervention struct {
	Reason string `bson:"reason" json:"reason"`

	Code CaseEmergencyCode `bson:"code" json:"code"`

	Address Address `bson:"address" json:"address"`
}
