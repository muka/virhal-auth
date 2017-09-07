package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"gitlab.fbk.eu/essence/essence-auth/api"
)

func main() {

	log.SetLevel(log.DebugLevel)

	err := api.StartDefault()
	if err != nil {
		log.Fatalf("Failed to setup: %s", err.Error())
		os.Exit(1)
	}
}
