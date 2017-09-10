package acl

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/casbin/casbin"
	mongodbadapter "github.com/casbin/mongodb-adapter"
	"github.com/spf13/viper"
)

var enforcer *casbin.Enforcer

func Setup() error {
	log.Debug("Initializing ACL")
	a := mongodbadapter.NewAdapter(strings.Join(viper.GetStringSlice("mongodb"), ","))
	enforcer = casbin.NewEnforcer(viper.GetString("acl_model"), a)
	err := enforcer.LoadPolicy()
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	enforcer = nil
	return nil
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
