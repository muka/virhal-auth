package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.fbk.eu/essence/essence-auth/db"
	"gitlab.fbk.eu/essence/essence-auth/model"
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

	u := model.NewUser()
	err := c.BindJSON(&u)
	if err != nil {
		panic(err)
	}

	err = db.UserCreate(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register"})
	}

	c.JSON(http.StatusAccepted, u)
}
