package casefile

type Address struct {
	County string `bson:"county" json:"county"`
	City   string `bson:"city" json:"city"`
	Street string `bson:"street" json:"street"`
	Nr     string `bson:"nr" json:"nr"`
	Bl     string `bson:"bl" json:"bl"`
	Sc     string `bson:"sc" json:"sc"`
	Et     string `bson:"et" json:"et"`
	Ap     string `bson:"ap" json:"ap"`
}
