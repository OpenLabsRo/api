package casefile

type CasePatient struct {
	FirstName   string  `bson:"firstName" json:"firstName"`
	LastName    string  `bson:"lastName" json:"lastName"`
	Age         int     `bson:"age" json:"age"`
	CNP         string  `bson:"cnp" json:"cnp"`
	Sex         string  `bson:"sex" json:"sex"`
	HomeAddress Address `bson:"homeAddress" json:"homeAddress"`
}
