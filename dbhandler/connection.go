package dbhandler

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	//TODO: import data via env
	//ConnectionString = "mongodb://mongo:27017"
	ConnectionString = "mongodb://localhost:27017"
	DatabaseName     = "userdb"
	UserCollection   = "users"
)

var Client *mongo.Client

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//ping db to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	fmt.Println("Successfully connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := Client.Database(DatabaseName).Collection(collectionName)
	if collection == nil {
		log.Fatal("Collection couldn't be found")
	}
	return collection
}
