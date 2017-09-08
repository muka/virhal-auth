package db

import (
	"gopkg.in/mgo.v2"
)

//IndexDefinition groups index definitions per collection
type IndexDefinition struct {
	coll    string
	indexes []mgo.Index
}

func getIndexes() []IndexDefinition {
	return []IndexDefinition{
		// User indexes
		IndexDefinition{
			coll: UserCollection,
			indexes: []mgo.Index{
				mgo.Index{
					Key:        []string{"username"},
					Unique:     true,
					DropDups:   true,
					Background: true,
					Sparse:     true,
				},
				mgo.Index{
					Key:        []string{"email"},
					Unique:     true,
					DropDups:   true,
					Background: true,
					Sparse:     true,
				},
			},
		},
		// Token indexes
		IndexDefinition{
			coll: TokenCollection,
			indexes: []mgo.Index{
				mgo.Index{
					Key:        []string{"token"},
					Unique:     true,
					DropDups:   true,
					Background: true,
					Sparse:     true,
				},
			},
		},
	}
}
