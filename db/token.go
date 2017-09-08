package db

import (
	log "github.com/Sirupsen/logrus"
	"gitlab.fbk.eu/essence/essence-auth/errors"
	"gitlab.fbk.eu/essence/essence-auth/model"
	"gopkg.in/mgo.v2"
)

//TokenColl return token collection
func TokenColl() *mgo.Collection {
	return Collection(TokenCollection)
}

//TokenCreate create a token
func TokenCreate() (model.Token, *errors.APIError) {

	t := model.NewToken()

	err := t.GenerateJWTToken()
	if err != nil {
		return t, errors.InternalServerError()
	}

	err = TokenColl().Insert(t)
	if err != nil {
		log.Errorf("Failed to store token: %s", err.Error())
		return t, errors.InternalServerError()
	}

	return t, nil
}

//TokenUpdate update an user
func TokenUpdate(t *model.Token, regenerateToken bool) *errors.APIError {
	if regenerateToken {
		err := t.GenerateJWTToken()
		if err != nil {
			return errors.InternalServerError()
		}
	}
	err := TokenColl().UpdateId(t.ObjectID, t)
	if err != nil {
		log.Errorf("Failed to update token: %s", err.Error())
		return errors.InternalServerError()
	}
	return nil
}

//TokenDelete delete an user
func TokenDelete(t model.Token) *errors.APIError {
	err := TokenColl().RemoveId(t.ObjectID)
	if err != nil {
		log.Errorf("Failed to delete token: %s", err.Error())
		return errors.InternalServerError()
	}
	return nil
}
