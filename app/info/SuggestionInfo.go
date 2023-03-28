package info

import (
	"github.com/vinllen/mgo/bson"
)

// 建议
type Suggestion struct {
	Id         bson.ObjectId `bson:"_id"`
	UserId     bson.ObjectId `UserId`
	Addr       string        `Addr`
	Suggestion string        `Suggestion`
}
