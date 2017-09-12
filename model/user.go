package model

import (
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

//Profile for user related generic data
type Profile map[string]interface{}

// User profile informations
type User struct {
	ObjectID     bson.ObjectId `json:"-" bson:"_id"`
	ID           string        `json:"id"`
	Enabled      bool          `json:"enabled"`
	Username     string        `json:"username" binding:"required,min=3,max=64"`
	Password     string        `json:"password,omitempty" binding:"required,min=3"`
	Email        string        `json:"email" binding:"required,min=4,email"`
	Roles        []Role        `json:"roles" binding:"required"`
	SessionToken string        `json:"sessionToken,omitempty"`
	Applications []string      `json:"applications,omitempty"`
	Profile      Profile       `json:"profile,omitempty"`
}

//ToPublicUser return a specialized struct for serialization
func (u *User) ToPublicUser() PublicUser {
	p := PublicUser{
		User: u,
	}
	return p
}

type omit *struct{}

//PublicUser expose defined fields on JSON marshalling
type PublicUser struct {
	*User
	ObjectID omit `json:"-"`
	Password omit `json:"password,omitempty"`
}

//NewUser init an user
func NewUser() User {
	return User{
		Roles:    make([]Role, 0),
		ID:       uuid.NewV4().String(),
		ObjectID: bson.NewObjectId(),
		Profile:  Profile{},
	}
}
