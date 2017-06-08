package model

type db interface {
	GetAllUsers() ([]*User, error)
	GetUser(login, password string) (*User, error)
	SetUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(u *User) error

	GetAllProducts() ([]*Product, error)
	GetProduct(productName string) (*Product, error)
	SetProduct(p *Product) error
	UpdateProduct(p *Product) error
	DeleteProduct(u *Product) error
}
