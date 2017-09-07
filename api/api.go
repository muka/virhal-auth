package api

import (
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"gitlab.fbk.eu/essence/essence-auth/db"
)

var listener net.Listener

const (
	defaultDatabase = "essence"
	defaultDbAddr   = "mongo"
	defaultPort     = ":8000"
)

func dbDefaults() *db.State {
	return &db.State{
		Addrs:    []string{defaultDbAddr},
		Database: defaultDatabase,
	}
}

//SetupDefault init api
func SetupDefault() error {
	return Setup(dbDefaults())
}

//Setup init api
func Setup(state *db.State) error {
	log.Debugf("Connecting to DB %s@%+v", state.Database, state.Addrs)
	err := db.Connect(state)
	if err != nil {
		return err
	}
	return nil
}

//StartDefault the server with default config
func StartDefault() error {
	return Start(defaultPort, dbDefaults())
}

//Start the server
func Start(address string, db *db.State) error {

	err := Setup(db)
	if err != nil {
		return err
	}

	log.Debug("Init router")
	r := gin.Default()

	r.POST("/register", UserRegister)

	err = manners.ListenAndServe(address, r)
	if err != nil {
		return err
	}

	log.Debug("Listening on %s", address)
	return nil
}

//Stop the server
func Stop() error {
	manners.Close()
	return nil
}
