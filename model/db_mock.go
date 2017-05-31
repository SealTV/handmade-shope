package model

import (
	"time"
)

type DBMock struct {
	users    []*User
	products []*Product
}

func (db *DBMock) GetAllUsers() ([]*User, error) {
	return db.users, nil
}

func (db *DBMock) GetUser(login, password string) (User, error) {
	for _, uu := range db.users {
		if uu.UserName == login && uu.Password == password {
			return *uu, nil
		}
	}
	return User{}, nil
}

func (db *DBMock) SetUser(u User) error {
	db.users = append(db.users, &u)
	return nil
}

func (db *DBMock) UpdateUser(u User) error {
	for _, uu := range db.users {
		if uu.ID == u.ID {
			uu.ID = u.ID
			uu.UserName = u.UserName
			uu.Email = u.Email
			uu.Password = u.Password
			return nil
		}
	}
	return nil
}

func (db *DBMock) GetAllProducts() ([]*Product, error) {
	return db.products, nil
}

func (db *DBMock) GetProduct(productName string) (Product, error) {
	for _, p := range db.products {
		if p.Name == productName {
			return *p, nil
		}
	}
	return Product{}, nil
}

func (db *DBMock) SetProduct(p Product) error {
	db.products = append(db.products, &p)
	return nil
}

func (db *DBMock) UpdateProduct(p Product) error {
	for _, pp := range db.products {
		if p.ID == p.ID {
			pp.Name = p.Name
			pp.Desription = p.Desription
			pp.Image = p.Image
			pp.Price = p.Price
			pp.UpdatedOn = time.Now()
		}
	}
	return nil
}
