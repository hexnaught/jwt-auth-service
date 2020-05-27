package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jokerdan/jwt-auth-service/model"
)

// registerHandler handles the creation of a user
func registerHandler(c *gin.Context) {
	user := &model.User{}

	err := c.BindJSON(user)
	if err != nil {
		log.Printf("Error: %+v", c.Errors)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateUserCreate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer ctxCancel()

	hashedPassword, err := hashString(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user.Password = hashedPassword

	result, err := db.Users.InsertOne(
		ctx,
		user,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = result.InsertedID

	// At this point, the user has registered and all is well, lets also auth
	// them and inject the token in to the response
	token := generateToken(user)

	c.JSON(
		http.StatusCreated,
		token,
	)
}
