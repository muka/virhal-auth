package api

import (
	"context"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.fbk.eu/essence/essence-auth/acl"
	"gitlab.fbk.eu/essence/essence-auth/db"
)

var server *http.Server

const appPrefix = "auth"

//LoadConfiguration load configurations
func LoadConfiguration() error {

	viper.SetDefault("log_level", "info")
	viper.SetDefault("mongodb", []string{"127.0.0.1"})
	viper.SetDefault("database", "auth_api")
	viper.SetDefault("listen", ":8000")
	viper.SetDefault("acl_model", "./acl_model.conf")

	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetEnvPrefix(appPrefix)
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Errorf("Failed to read configuration file: %s", err.Error())
		return err
	}
	return nil
}

//Setup init api
func Setup() error {

	log.Warnf("cfg %+v", viper.AllSettings())

	logLevel := log.InfoLevel
	switch viper.GetString("log_level") {
	case "debug":
		logLevel = log.DebugLevel
		break
	case "warn":
		logLevel = log.WarnLevel
		break
	case "error":
		logLevel = log.ErrorLevel
		break
	case "fatal":
		logLevel = log.FatalLevel
		break
	case "panic":
		logLevel = log.PanicLevel
		break
	}
	log.SetLevel(logLevel)

	log.Debug("Starting setup")
	err := db.Connect()
	if err != nil {
		return err
	}

	err = acl.Setup()
	if err != nil {
		return err
	}

	return nil
}

//Start the server
func Start() error {

	err := Setup()
	if err != nil {
		return err
	}

	r := registerRoutes()

	addr := viper.GetString("listen")
	server = &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			log.Warnf("Failed to start server %s", err.Error())
			return
		}
	}()

	log.Debugf("Listening on http://%s", addr)
	return nil
}

//Stop the server
func Stop() error {
	db.Disconnect()
	acl.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
