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

	viper.Set("log_level", "debug")
	viper.Set("listen", testPort)
	viper.Set("database", "auth_test")

	log.Debugf("Listening on %s", viper.Get("listen"))

	err = api.Start()
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Millisecond * 500)

	retCode := m.Run()

	if err := api.Stop(); err != nil {
		panic(err)
	}
	os.Exit(retCode)
}

func registerUser(u model.RequestRegister) (*model.User, error) {
	uri := fmt.Sprintf("http://localhost%s/%s", testPort, "register")
	user := model.NewUser()
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(u).
		SetResult(&user). // or SetResult(AuthSuccess{}).
		Post(uri)

	log.Printf("RAW %+v", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func loginUser(l model.RequestLogin) (model.ResponseLogin, error) {

	lr := model.ResponseLogin{}
	uri := fmt.Sprintf("http://localhost%s/%s", testPort, "login")
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(l).
		SetResult(&lr).
		Post(uri)

	log.Printf("RAW %+v", string(resp.Body()))

	return lr, err
}

func TestUserRegister(t *testing.T) {

	u := model.RequestRegister{
		Username: "test" + strconv.Itoa(int(time.Now().UnixNano())),
		Password: "secret",
	}
	u.Email = u.Username + "@test.local"

	_, err := registerUser(u)
	if err != nil {
		t.Fatal(err)
	}

}

func TestUserLogin(t *testing.T) {

	u := model.RequestRegister{
		Username: "test" + strconv.Itoa(int(time.Now().UnixNano())),
		Password: "secret",
	}
	u.Email = u.Username + "@test.local"

	_, err := registerUser(u)
	if err != nil {
		t.Fatal(err)
	}

	l := model.RequestLogin{
		Username: u.Username,
		Password: u.Password,
	}

	_, err = loginUser(l)
	if err != nil {
		t.Fatal(err)
	}

}
