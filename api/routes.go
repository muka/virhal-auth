package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"gitlab.fbk.eu/essence/essence-auth/errors"
	"gitlab.fbk.eu/essence/essence-auth/model"
)

func registerRoutes() *gin.Engine {

	log.Debug("Init router")
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {

		u := model.RequestRegister{}
		err := c.BindJSON(&u)
		if err != nil {
			verr := errors.Validation(err)
			c.JSON(verr.Code, verr)
			return
		}

		user, apiErr := UserRegister(u)
		if apiErr != nil {
			c.JSON(apiErr.Code, apiErr)
			return
		}

		c.JSON(http.StatusAccepted, user.ToPublicUser())
	})

	r.POST("/login", func(c *gin.Context) {

		r := model.RequestLogin{}
		err := c.BindJSON(&r)
		if err != nil {
			verr := errors.Validation(err)
			c.JSON(verr.Code, verr)
			return
		}

		u, sso, e := UserLogin(r)
		if e != nil {
			c.JSON(e.Code, e)
			return
		}

		res := model.NewResponseLogin(&u, sso)
		c.JSON(http.StatusAccepted, res)

	})

	auth := r.Group("/", AuthHandler)
	auth.POST("/authorized", func(c *gin.Context) {
		//
		// TODO fetch body content
		//
		res, err := IsAuthorized()
		if err != nil {
			c.JSON(err.Code, err)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	return r
}
