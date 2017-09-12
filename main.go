package main

import (
	"os"
	"os/signal"

	log "github.com/Sirupsen/logrus"
	"github.com/muka/virhal-auth/api"
)

func main() {

	err := api.LoadConfiguration()
	if err != nil {
		log.Fatalf("Failed to configure: %s", err.Error())
		os.Exit(1)
	}

	err = api.Start()
	if err != nil {
		log.Fatalf("Failed to start: %s", err.Error())
		os.Exit(1)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	err = api.Stop()
	if err != nil {
		log.Fatalf("Failed to close: %s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
