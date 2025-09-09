package dbhandler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitDB() {
	// make env vars accessable
	loadEnv()
	// create DB URI
	connectionString := createDBConnectionURI()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
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
	collection := Client.Database(getDBName()).Collection(collectionName)
	if collection == nil {
		log.Fatal("Collection couldn't be found")
	}
	return collection
}

func loadEnv() {
	err := godotenv.Load(os.Getenv("ENV_FILE"))
	if err != nil {
		log.Fatal("Error loading env file, '${ENV_FILE} Variable not set'")
	}
}

func createDBConnectionURI() string {
	result := fmt.Sprintf("%s://%s:%s", os.Getenv("DB_SCHEME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	return result
}

func getDBName() string {
	userDBName := os.Getenv("DB_NAME_USER")
	return userDBName
}

func GetUserCollection() string {
	userCollection := os.Getenv("DB_USER_COLLECTION")
	return userCollection
}
