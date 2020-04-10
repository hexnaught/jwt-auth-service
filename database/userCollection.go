package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *MongoHelper) setupUserCollection() {
	db.Users = db.Database.Collection("users")

	// User collection indexes
	userIndexModels := []mongo.IndexModel{
		{
			Keys: bson.M{
				"email": 1,
			}, Options: options.Index().SetUnique(true),
		}, {
			Keys: bson.M{
				"username": 1,
			}, Options: options.Index().SetUnique(true),
		},
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer ctxCancel()

	_, err := db.Users.Indexes().CreateMany(ctx, userIndexModels)
	if err != nil {
		log.Fatalf("creating user indexes failed: %+v", err.Error())
	}
}
