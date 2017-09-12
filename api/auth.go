package api

import (
	goerrors "errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/muka/virhal-auth/errors"
	"strings"
)

const defaultTokenHeader = "Authorization"
const defaultTokenPrefix = "Authorization"

//AuthHandler handle authentication
func AuthHandler(c *gin.Context) {
	log.Debug("Auth request")

	token, err := getHeaderToken(c, defaultTokenHeader)
	if err != nil {
		uerr := errors.Unauthorized()
		c.JSON(uerr.Code, uerr)
		return
	}

	log.Debugf("token %s", token)

	c.Next()
}

func getHeaderToken(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", goerrors.New("auth header empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == defaultTokenPrefix) {
		return "", fmt.Errorf("Invalid auth header, missing %s", defaultTokenPrefix)
	}

	return parts[1], nil
}

//IsAuthorized check for authorization
func IsAuthorized() (bool, *errors.APIError) {
	return false, nil
}
