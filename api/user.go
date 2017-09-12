package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/muka/virhal-auth/db"
	"github.com/muka/virhal-auth/errors"
	"github.com/muka/virhal-auth/model"
	"golang.org/x/crypto/bcrypt"
)

func decodeUser(r *http.Request) (model.User, error) {
	u := model.NewUser()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

//UserRegister register a new user
func UserRegister(u model.RequestRegister) (model.User, *errors.APIError) {

	user := u.ToUser()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("Failed to hash password: %s", err.Error())
		err1 := errors.InternalServerError()
		return user, err1
	}

	user.Password = string(hashedPassword)

	e := db.UserCreate(user)
	if e != nil {
		return user, e
	}

	return user, nil
}

//UserLogin login a user
func UserLogin(r model.RequestLogin) (model.User, *errors.APIError) {

	user, err := db.UserLogin(r.Username, r.Password)

	if err != nil {
		return user, err
	}

	log.Debug("Create login token")
	tokens, terr := db.TokenCreate(&user, 1)
	if terr != nil {
		log.Debug("Failed to create JWT tokens")
		return user, terr
	}
	user.SessionToken = tokens[0].Value

	return user, nil
}
