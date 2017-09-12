package model

import (
	"errors"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

const defaultExpires = time.Minute * 30

//Token for authentication
type Token struct {
	ObjectID bson.ObjectId `json:"-" bson:"_id"`
	UserID   string        `json:"userId"`
	Value    string        `json:"token"`
	Expires  time.Time     `json:"expires"`
	Created  time.Time
}

//GenerateJWTToken generate a JWT token
func (t *Token) GenerateJWTToken() error {

	secret := []byte(viper.GetString("jwt_token"))
	if len(secret) == 0 {
		return errors.New("jwt_token must be set")
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
		Created:  time.Now(),
		Expires:  time.Now().Add(defaultExpires),
		ObjectID: bson.NewObjectId(),
		UserID:   user.ID,
	}
}
