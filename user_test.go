package main

import (
	"fmt"
	"strconv"

	"os"
	"testing"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-resty/resty"
	"github.com/spf13/viper"
	"gitlab.fbk.eu/essence/essence-auth/api"
	"gitlab.fbk.eu/essence/essence-auth/model"
)

const testPort = ":8009"

func TestMain(m *testing.M) {

	log.SetLevel(log.DebugLevel)

	err := api.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	viper.Set("listen", testPort)
	viper.Set("database", "auth_test")

	go func() {
		err := api.Start()
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

	u := model.RequestRegister{
		Username: "test" + strconv.Itoa(int(time.Now().UnixNano())),
		Password: "secret",
	}
	u.Email = u.Username + "@test.local"

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
