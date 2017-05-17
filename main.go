package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	const (
		Host     = "shopdb.documents.azure.com:10255"
		Username = "shopdb"
		Password = "beyowzrb8wZytMpIQ50IvWTzfOZaSstSLtAloPZyiy2Jv0oUN5CZqs59YQEV73Mzz0QkZKqNmqZ9mgSXUJy1kw=="
		Database = "test"
	)

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{Host},
		Username: Username,
		Password: Password,
		Database: Database,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Printf("Connected to %v!\n", session.LiveServers())

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
