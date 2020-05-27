package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jokerdan/jwt-auth-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// loginHandler is a handler for the API login endpoint
func loginHandler(c *gin.Context) {
	user := &model.User{}
	fetchedUser := &model.User{}

	if c.BindJSON(user) != nil {
		log.Printf("Error: %+v", c.Errors)
		return
	}

	err := user.ValidateUserLogin()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer ctxCancel()

	err = db.Users.FindOne(ctx, bson.D{{Key: "username", Value: user.Username}}).Decode(&fetchedUser)
	if err != nil {
		// Likely, no user found
		c.JSON(http.StatusNotFound, gin.H{"error": "no user found for supplied details"})
		return
	}

	hasAuthed := verifyPassword(fetchedUser.Password, user.Password)
	if !hasAuthed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}

	// If the user has been able to auth, lets double check that their password meets the hash cost
	// if it isn't let us rehash the password with the adjusted cost
	checkAndUpdateUserPasswordHash(fetchedUser, user.Password)

	token := generateToken(fetchedUser)

	c.JSON(
		http.StatusOK,
		token,
	)
}

// generateToken is a wrapper for generating the token and nulling the 'password' field
// we were sent so that it doesn't get returned back to the calling service
func generateToken(user *model.User) string {
	token, err := createToken(user.ID.(primitive.ObjectID).Hex(), user.Username, user.Email)
	if err != nil {
		return ""
	}

	return token
}

// checkAndUpdateUserPasswordHash checks the bcrypt cost factor user for the users hashed password after they have
// successfully authed, if it is lower than the apps currently used one, re-hash and re-save the password
func checkAndUpdateUserPasswordHash(fetchedUser *model.User, password string) {
	ctx, ctxCancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer ctxCancel()

	cost, err := bcrypt.Cost([]byte(fetchedUser.Password))
	if err != nil {
		log.Printf("error: failed to get cost factor of hashed password for user %v", fetchedUser.ID)
	}

	if cost < appConfig.BCryptCostFactor {
		hashedPassword, err := hashString(password)
		if err != nil {
			log.Printf("error: failed to update bcrypt cost and rehash password for user %v", fetchedUser.ID)
			return
		}
		db.Users.UpdateOne(
			ctx,
			bson.M{"_id": fetchedUser.ID.(primitive.ObjectID)},
			bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: hashedPassword}}}},
		)
	}
}
