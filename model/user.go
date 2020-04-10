package model

import "errors"

// User represents what our concept of a user is
type User struct {
	ID       interface{} `bson:"_id,omitempty"      json:"_id,omitempty"`
	Username string      `bson:"username"           json:"username"`
	Email    string      `bson:"email,omitempty"    json:"email"`
	Password string      `bson:"password,omitempty" json:"password,omitempty"`
	Token    string      `bson:"token,omitempty"    json:"token,omitempty"`
}

// ValidateUserCreate validates the user at point of creation
func (u User) ValidateUserCreate() error {
	if u.Username == "" && u.Email == "" && u.Password == "" {
		return errors.New("a username, email and password must be supplied to register")
	}
	return nil
}

// ValidateUserLogin validates that the requried credentials to log in have been sent
func (u User) ValidateUserLogin() error {
	if u.Password == "" || (u.Username == "" && u.Email == "") {
		return errors.New("a username or email need to be provided with a password to login")
	}
	return nil
}
