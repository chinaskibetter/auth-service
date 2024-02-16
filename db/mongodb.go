package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	mongoURI       = "mongodb://localhost:27017"
	dbName         = "auth"
	collectionName = "tokens"
)

var mongoClient *mongo.Client
var collection *mongo.Collection

func InitMongoClient() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	mongoClient = client
	collection = client.Database(dbName).Collection(collectionName)
	return nil
}

func GetCollection() *mongo.Collection {
	return collection
}
