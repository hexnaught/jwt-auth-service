package api

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func verifyPassword(hashedPassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("password validation error: %+v", err)
		return false
	}
	return true
}

func hashString(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), appConfig.BCryptCostFactor)
	if err != nil {
		log.Printf("error hashing password: %+v", err)
		return "", err
	}
	return string(passwordHash), nil
}

func createToken(userID, username, email string) (string, error) {
	var err error

	jwtClaims := jwt.MapClaims{
		"iat":      time.Now().Unix(),
		"user_id":  userID,
		"username": username,
	}

	if appConfig.JWTTTL != 0 {
		jwtClaims["exp"] = time.Now().Add(time.Minute * time.Duration(appConfig.JWTTTL)).Unix()
	}

	if appConfig.JWTISS != "" {
		jwtClaims["iss"] = appConfig.JWTISS
	}

	if len(appConfig.JWTAUD) != 0 {
		jwtClaims["aud"] = appConfig.JWTAUD
	}

	newJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	token, err := newJWT.SignedString([]byte(appConfig.JWTSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func validateToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected jwt signing method: %v", token.Header["alg"])
		}
		return []byte(appConfig.JWTSecret), nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
