package db

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
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
)

var state *State

//State store db state and config
type State struct {
	session  *mgo.Session
	Addrs    []string
	Database string
}

//Connect to database
func Connect(cstate *State) error {

	state = cstate

	var err error
	state.session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    state.Addrs,
		Database: state.Database,
	})

	if err != nil {
		return err
	}

	// Optional. Switch the session to a monotonic behavior.
	state.session.SetMode(mgo.Monotonic, true)

	log.Debug("Session initialized")
	return nil
}

//Disconnect from database
func Disconnect() {
	state.session.Close()
}

//Collection load a collection
func Collection(coll string) *mgo.Collection {
	log.Debugf("%+v", state.session)
	if state.session == nil {
		panic(fmt.Errorf("Cannot get session from %+v", state.Addrs))
	}
	d := state.session.DB(state.Database)
	if d == nil {
		panic(fmt.Errorf("Cannot get Database %s", state.Database))
	}
	c := d.C(coll)
	if c == nil {
		panic(fmt.Errorf("Cannot get Collection %s", coll))
	}
	return c
}
