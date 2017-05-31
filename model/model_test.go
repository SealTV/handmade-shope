package model

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestNew(t *testing.T) {
	var db DBMock
	db.SetUser(User{ID: bson.NewObjectId(), UserName: "user1", Email: "email1@example.com", Password: "pass"})
	db.SetUser(User{ID: bson.NewObjectId(), UserName: "user2", Email: "email2@example.com", Password: "pass"})
	db.SetUser(User{ID: bson.NewObjectId(), UserName: "user3", Email: "email3@example.com", Password: "pass"})
	m := New(&db)

	if u, _ := m.Users(); len(u) != 3 {
		t.Errorf("Users() len = %d", len(u))
	}
}

func TestUsers(t *testing.T) {
	var db DBMock
	db.SetProduct(Product{ID: bson.NewObjectId(), Name: "name1", Desription: "Desription1", Price: 1})
	db.SetProduct(Product{ID: bson.NewObjectId(), Name: "name2", Desription: "Desription2", Price: 2})
	db.SetProduct(Product{ID: bson.NewObjectId(), Name: "name3", Desription: "Desription3", Price: 3})
	m := New(&db)

	if p, _ := m.Products(); len(p) != 3 {
		t.Errorf("Products() len = %d", len(p))
	}
}
