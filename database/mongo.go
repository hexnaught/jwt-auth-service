package database

import (
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoHelper is a helper struct for interfacing with Mongo
type MongoHelper struct {
	Database *mongo.Database
	Users    *mongo.Collection
}

// NewMongoHelper creates a new instance of MongoHelper to make accessing
// the mongo DB easier
func NewMongoHelper() *MongoHelper {
	// Setup mongo client
	mongoHelper := &MongoHelper{}

	dbConfig := &MongoConfig{
		Host:         os.Getenv("MONGO_HOST"),
		Port:         os.Getenv("MONGO_PORT"),
		Username:     os.Getenv("MONGO_USERNAME"),
		Password:     os.Getenv("MONGO_PASSWORD"),
		AuthSource:   os.Getenv("MONGO_AUTH_SOURCE"),
		DatabaseName: os.Getenv("MONGO_DB_NAME"),
	}

	dbClient, err := setupMongoClient(dbConfig)

	if err != nil {
		log.Fatal("Error initialising client database")
	}
	err = dbClient.Connect(nil)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	mongoHelper.Database = dbClient.Database(dbConfig.DatabaseName)

	mongoHelper.setupUserCollection()

	return mongoHelper
}

// NewMongoClient instantiates a new Mongo client
func setupMongoClient(config *MongoConfig) (*mongo.Client, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%s/%s", config.Host, config.Port, config.DatabaseName)

	clientOptions := options.Client().ApplyURI(connectionString)
	// If this is not running against localhost mongo instance, try and connect with auth user creds
	if config.Host != "127.0.0.1" && config.Host != "localhost" && config.Host != "mongodb" {
		clientOptions.SetAuth(
			options.Credential{
				Username:   config.Username,
				Password:   config.Password,
				AuthSource: config.AuthSource,
			},
		)
	}

	client, err := mongo.NewClient(clientOptions)
	return client, err
}
