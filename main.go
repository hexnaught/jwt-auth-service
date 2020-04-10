package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jokerdan/jwt-auth-service/api"
	"github.com/jokerdan/jwt-auth-service/config"
	"github.com/jokerdan/jwt-auth-service/database"
)

func main() {

	appConfig := config.LoadAppConfig()

	mongoDB := database.NewMongoHelper()
	defer mongoDB.Database.Client().Disconnect(context.TODO())

	router := api.SetUp(mongoDB, appConfig)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(
		fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), port),
	)
}
