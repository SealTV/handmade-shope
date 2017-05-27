package db

import (
	"crypto/tls"
	"fmt"
	"net"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/SealTV/handmade-shope/model"
)

const (
	database = "Shop"
	products = "products"
	users    = "users"
)

// MongoConfig - config params for connection
type MongoConfig struct {
	Host     string
	Username string
	Password string
}

// MongoDB - coneection struct
type MongoDB struct {
	session  *mgo.Session
	users    *mgo.Collection
	products *mgo.Collection
}

// InitMongoDb - initialize MongoDB client
func InitMongoDb(cfg MongoConfig) (*MongoDB, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{cfg.Host},
		Username: cfg.Username,
		Password: cfg.Password,
		Database: database,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	})

	if err != nil {
		panic(err)
	}

	p := &MongoDB{
		session:  session,
		users:    session.DB(database).C(users),
		products: session.DB(database).C(products),
	}

	fmt.Printf("Connected to %v!\n", session.LiveServers())

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return p, nil
}

func (m *MongoDB) Close() {
	m.session.Close()
}

func (m *MongoDB) GetAllUsers() ([]*model.User, error) {
	result := []*model.User{}
	err := m.users.Find(bson.M{}).All(&result)
	return result, err
}

func (m *MongoDB) GetUser(login, password string) (model.User, error) {
	result := model.User{}
	err := m.users.Find(bson.M{"login": login, "password": password}).One(&result)
	return result, err
}

func (m *MongoDB) SetUser(u model.User) error {
	err := m.users.Insert(&u)
	return err
}

func (m *MongoDB) UpdateUser(u model.User) error {
	err := m.users.Update(bson.M{"login": u.Email}, u)
	return err
}

func (m *MongoDB) GetAllProducts() ([]*model.Product, error) {
	result := []*model.Product{}
	err := m.products.Find(bson.M{}).All(&result)
	return result, err
}

func (m *MongoDB) GetProduct(productName string) (model.Product, error) {
	result := model.Product{}
	err := m.products.Find(bson.M{"name": productName}).One(&result)
	return result, err
}

func (m *MongoDB) SetProduct(p model.Product) error {
	err := m.products.Insert(&p)
	return err
}

func (m *MongoDB) UpdateProduct(p model.Product) error {
	err := m.products.Update(bson.M{"name": p.Name}, p)
	return err
}
