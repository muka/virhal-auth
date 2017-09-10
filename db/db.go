package db

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

const (
	//UserCollection collection for users
	UserCollection = "users"
	//RoleCollection collection for roles
	RoleCollection = "roles"
	//PermissionCollection collection for permissions
	PermissionCollection = "permissions"
	//ServiceCollection collection for services
	ServiceCollection = "services"
	//TokenCollection collection for tokens
	TokenCollection = "tokens"
)

var session *mgo.Session

//Connect to database
func Connect() error {

	if session != nil {
		if err := session.Ping(); err == nil {
			return nil
		}
		session.Close()
		session = nil
	}

	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    viper.GetStringSlice("mongodb"),
		Database: viper.GetString("database"),
	})

	if err != nil {
		return err
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	err = setupIndexes()
	if err != nil {
		return err
	}

	log.Debug("Session initialized")
	return nil
}

func setupIndexes() error {
	log.Debug("Setup indexes")
	defs := getIndexes()
	for d := 0; d < len(defs); d++ {
		def := defs[d]
		for i := 0; i < len(def.indexes); i++ {
			index := def.indexes[i]
			err := Collection(def.coll).EnsureIndex(index)
			if err != nil {
				return err
			}
		}
	}
	log.Debug("Index setup done")
	return nil
}

//Disconnect from database
func Disconnect() {
	session.Close()
}

//Collection load a collection
func Collection(coll string) *mgo.Collection {

	addrs := viper.GetStringSlice("mongodb")
	db := viper.GetString("database")

	if session == nil {
		panic(fmt.Errorf("Cannot get session from %+v", addrs))
	}
	d := session.DB(db)
	if d == nil {
		panic(fmt.Errorf("Cannot get Database %s", db))
	}
	c := d.C(coll)
	if c == nil {
		panic(fmt.Errorf("Cannot get Collection %s", coll))
	}
	return c
}
