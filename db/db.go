package db

import (
	"api/env"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.Background()
var Client *mongo.Client

var Accounts *mongo.Collection
var Dispatchers *mongo.Collection
var Paramedics *mongo.Collection

var Cases *mongo.Collection

var Teams *mongo.Collection
var Ambulances *mongo.Collection

var Hospitals *mongo.Collection

var Events *mongo.Collection

func InitDB() (err error) {
	Client, err = mongo.Connect(
		Ctx,
		options.Client().ApplyURI(env.MONGO_URI),
	)

	if err != nil {
		return
	}

	Accounts = GetCollection("accounts", Client)
	Dispatchers = GetCollection("dispatchers", Client)
	Paramedics = GetCollection("paramedics", Client)

	Cases = GetCollection("cases", Client)

	Teams = GetCollection("teams", Client)
	Ambulances = GetCollection("ambulances", Client)

	Hospitals = GetCollection("hospitals", Client)

	Events = GetCollection("events", Client)

	fmt.Println("Connected to MongoDB")
	return nil
}

func GetCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	return client.Database("dev").Collection(collectionName)
}
