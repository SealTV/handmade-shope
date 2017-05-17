package models

import "time"

type User struct {
	UserName string `bson:"userName"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type Product struct {
	Name       string
	Desription string
	Image      []byte
	Price      int
	CreatedOn  time.Time
	UpdatedOn  time.Time
}
