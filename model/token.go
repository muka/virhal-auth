package model

import (
	"errors"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

const defaultExpires = time.Minute * 30

//Token for authentication
type Token struct {
	ObjectID bson.ObjectId `json:"_id,omitempty"`
	UserID   string        `json:"UserID"`
	Expires  time.Time     `json:"expires"`
	Created  time.Time
	Value    string `json:"token"`
}

//GenerateJWTToken generate a JWT token
func (t *Token) GenerateJWTToken() error {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic(errors.New("JWT_SECRET not set"))
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":      t.Created.Unix(),
		"orig_iat": t.Created.Unix(),
		"exp":      t.Expires.Unix(),
		"ID":       t.UserID,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Errorf("Failed to generate token: %s", err.Error())
		return err
	}

	t.Value = tokenString

	return nil
}

//NewToken create a new token
func NewToken(user *User) Token {
	return Token{
		Expires: time.Now().Add(defaultExpires),
		Created: time.Now(),
		UserID:  user.UserID,
	}
}
