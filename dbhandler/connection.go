package dbhandler

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required"`
	Phone     string             `json:"phone" bson:"phone,omitempty"`
	Username  string             `json:"username" bson:"username" binding:"required,min=1,max=30"`
	Password  string             `json:"passwortd" bson:"password" bindign:"required,min=15"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

const (
	ConnectionString = "mongodb://mongo:27017"
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
