package main

import (
	"fmt"

	"os"
	"testing"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-resty/resty"
	"gitlab.fbk.eu/essence/essence-auth/api"
	"gitlab.fbk.eu/essence/essence-auth/db"
	"gitlab.fbk.eu/essence/essence-auth/model"
)

const testPort = ":8009"

func TestMain(m *testing.M) {

	log.SetLevel(log.DebugLevel)

	go func() {
		err := api.Start(testPort, &db.State{
			Addrs:    []string{"mongo"},
			Database: "essence_test",
		})
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second * 2)

	retCode := m.Run()

	if err := api.Stop(); err != nil {
		panic(err)
	}
	os.Exit(retCode)
}

func TestUserRegister(t *testing.T) {

	u := model.NewUser()
	uri := fmt.Sprintf("http://localhost%s/%s", testPort, "register")
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(u).
		SetResult(&model.User{}). // or SetResult(AuthSuccess{}).
		Post(uri)

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", resp)

}
