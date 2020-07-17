package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hexnaught/jwt-auth-service/config"
	"github.com/hexnaught/jwt-auth-service/database"
)

var db *database.MongoHelper
var appConfig *config.App

// SetUp creates a new instance of a gin.Engine with our endpoints registered
func SetUp(mongoDBHelper *database.MongoHelper, App *config.App) *gin.Engine {

	db = mongoDBHelper
	appConfig = App

	router := gin.Default()

	// ? Do I want to use logrus
	// log := logrus.New()
	// router := gin.New()
	// router.Use(ginlogrus.Logger(log), gin.Recovery())

	baseAPI := router.Group("/api/v1")

	userGroup := baseAPI.Group("/user")
	{
		userGroup.POST("/register", registerHandler)
		userGroup.POST("/login", loginHandler)
	}

	tokenGroup := baseAPI.Group("/token")
	{
		tokenGroup.GET("/validate", validateTokenHandler)
	}

	return router
}
