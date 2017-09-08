package api

import (
	"context"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"gitlab.fbk.eu/essence/essence-auth/db"
)

var server *http.Server

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
	r.POST("/login", UserLogin)

	auth := r.Group("/", AuthHandler)
	auth.POST("/authorized", IsAuthorized)

	server = &http.Server{
		Addr:    address,
		Handler: r,
	}

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			log.Warnf("Failed to start server %s", err.Error())
			return
		}
	}()

	log.Debugf("Listening on http://%s", address)
	return nil
}

//Stop the server
func Stop() error {
	db.Disconnect()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
