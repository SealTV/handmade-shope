package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product model
type Product struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Desription string        `bson:"description" json:"description"`
	Image      []byte        `bson:"image" json:"image"`
	Price      int           `bson:"price" json:"price"`
	CreatedOn  time.Time     `bson:"create_on" json:"create_on"`
	UpdatedOn  time.Time     `bson:"update_on" json:"update_on"`
}
