package db

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"gitlab.fbk.eu/essence/essence-auth/errors"
	"gitlab.fbk.eu/essence/essence-auth/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserColl user collection
func UserColl() *mgo.Collection {
	return Collection(UserCollection)
}

//UserCreate create a new user
func UserCreate(u model.User) *errors.APIError {

	err := UserColl().Insert(u)
	if err != nil {
		if mgo.IsDup(err) {
			if strings.Contains(err.Error(), "index: email_") {
				return errors.RegistrationFailed("Email already registered")
			}
			if strings.Contains(err.Error(), "index: username_") {
				return errors.RegistrationFailed("Username already registered")
			}
		}
		log.Errorf("Failed to insert user: %s", err.Error())
		return errors.InternalServerError()
	}

	return nil
}

//UserUpdate update an user
func UserUpdate(u model.User) *errors.APIError {
	err := UserColl().UpdateId(u.ObjectID, u)
	if err != nil {
		log.Errorf("User update failed: %s", err.Error())
		return errors.InternalServerError()
	}
	return nil
}

//UserDelete delete an user
func UserDelete(u model.User) *errors.APIError {
	err := UserColl().RemoveId(u.ObjectID)
	if err != nil {
		log.Errorf("User delete failed: %s", err.Error())
		return errors.InternalServerError()
	}
	return nil
}

//UserFind find users
func UserFind(query interface{}) []model.User {
	q := UserColl().Find(query)
	record := model.User{}
	list := make([]model.User, 0)
	for q.Iter().Next(record) {
		list = append(list, record)
	}
	return list
}

//UserLogin login a user
func UserLogin(username, password string) (model.User, *errors.APIError) {

	user := model.User{}
	if username == "" || password == "" {
		return user, errors.LoginFailed("Username and password are required")
	}

	err := UserColl().Find(bson.M{"username": username}).One(&user)
	if err != nil {
		log.Debugf("User not found: %s", username)
		return user, errors.LoginFailed("Invalid credentials")
	}

	log.Debugf("found user %s", user.Username)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			log.Debug("Password not matching")
			return user, errors.LoginFailed("Invalid credentials")
		}
		return user, errors.InternalServerError()
	}

	return user, nil
}
