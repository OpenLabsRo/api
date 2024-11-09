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

func InitDB() (err error) {
	Client, err = mongo.Connect(
		Ctx,
		options.Client().ApplyURI(env.MONGO_URI),
	)

	if err != nil {
		return
	}

	Accounts = GetCollection("accounts")
	Dispatchers = GetCollection("dispatchers")
	Paramedics = GetCollection("paramedics")
	Cases = GetCollection("cases")

	fmt.Println("Connected to MongoDB")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("dev").Collection(collectionName)
}
