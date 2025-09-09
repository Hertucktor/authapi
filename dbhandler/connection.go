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
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env.prod"
		log.Printf("ENV_FILE not set, using default: %s", envFile)
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading env file '%s': %v", envFile, err)
	}

	log.Printf("Successfully loaded environment from %s", envFile)
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
