package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"gitlab.fbk.eu/essence/essence-auth/db"
	"gitlab.fbk.eu/essence/essence-auth/errors"
	"gitlab.fbk.eu/essence/essence-auth/model"
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
func UserLogin(r model.RequestLogin) (model.User, model.Sso, *errors.APIError) {

	user, err := db.UserLogin(r.Username, r.Password)
	sso := model.Sso{}

	if err != nil {
		return user, sso, err
	}

	log.Debug("Create login token")
	tokens, terr := db.TokenCreate(&user, 9)
	if terr != nil {
		log.Debug("Failed to create JWT tokens")
		return user, sso, terr
	}
	user.SessionToken = tokens[0].Value
	//sso specific
	sso.Atlante = tokens[1].Value
	sso.Biophr = tokens[2].Value
	sso.Chino = tokens[3].Value
	sso.Cube3D = tokens[4].Value
	sso.FitForAll = tokens[5].Value
	sso.Raptor = tokens[6].Value
	sso.Trilogis = tokens[7].Value
	sso.Webrtc = tokens[8].Value

	return user, sso, nil
}
