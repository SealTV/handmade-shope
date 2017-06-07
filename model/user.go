package model

import "gopkg.in/mgo.v2/bson"

// User model
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Id       int64         `bson:"id,omitempty" json:"id"`
	UserName string        `bson:"login" json:"login"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
}
