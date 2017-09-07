package db

import (
	"gitlab.fbk.eu/essence/essence-auth/model"
	"gopkg.in/mgo.v2"
)

//UserColl user collection
func UserColl() *mgo.Collection {
	return Collection(UserCollection)
}

//UserCreate create a new user
func UserCreate(u model.User) error {
	return UserColl().Insert(u)
}

//UserUpdate update an user
func UserUpdate(u model.User) error {
	return UserColl().UpdateId(u.ID, u)
}

//UserDelete delete an user
func UserDelete(u model.User) error {
	return UserColl().RemoveId(u.ID)
}

//UserFind find users
func UserFind(query interface{}) []model.User {
	q := UserColl().Find(query)
	record := model.User{}
	list := make([]model.User, 0)
	for q.Iter().Next(record) {
		list = append(list, record)
	}
	return list
}
