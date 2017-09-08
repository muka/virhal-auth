package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
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
func UserRegister(c *gin.Context) {

	u := model.RequestRegister{}
	err := c.BindJSON(&u)
	if err != nil {
		verr := errors.Validation(err)
		c.JSON(verr.Code, verr)
		return
	}

	user := u.ToUser()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("Failed to hash password: %s", err.Error())
		err1 := errors.InternalServerError()
		c.JSON(err1.Code, err1)
		return
	}

	user.Password = string(hashedPassword)

	e := db.UserCreate(user)
	if e != nil {
		c.JSON(e.Code, e)
		return
	}

	c.JSON(http.StatusAccepted, user)
}

//UserLogin login a user
func UserLogin(c *gin.Context) {

	r := model.RequestLogin{}
	err := c.BindJSON(&r)
	if err != nil {
		verr := errors.Validation(err)
		c.JSON(verr.Code, verr)
		return
	}

	u, e := db.UserLogin(r.Username, r.Password)
	if e != nil {
		c.JSON(e.Code, e)
		return
	}

	res := model.NewResponseLogin(&u)
	c.JSON(http.StatusAccepted, res)
}
