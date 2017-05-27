package model

// User model
type User struct {
	UserName string `bson:"userName"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
