package model

import "gopkg.in/mgo.v2/bson"

// User model
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	UserName string        `bson:"userName" json:"login"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
}
