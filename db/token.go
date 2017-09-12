package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/muka/virhal-auth/errors"
	"github.com/muka/virhal-auth/model"
	"gopkg.in/mgo.v2"
)

//TokenColl return token collection
func TokenColl() *mgo.Collection {
	return Collection(TokenCollection)
}

//TokenCreate create JWT tokens
func TokenCreate(user *model.User, count int) ([]model.Token, *errors.APIError) {
	// required for []interface{} conversion
	itokens := make([]interface{}, count)
	tokens := make([]model.Token, count)
	for count > 0 {
		t := model.NewToken(user)
		err := t.GenerateJWTToken()
		if err != nil {
			return tokens, errors.InternalServerError()
		}
		count--
		tokens[count] = t
		itokens[count] = t
	}

	err := TokenColl().Insert(itokens...)
	if err != nil {
		log.Errorf("Failed to store token: %s", err.Error())
		return tokens, errors.InternalServerError()
	}

	return tokens, nil
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
