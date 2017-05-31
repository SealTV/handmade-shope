package model

// Model - base type
type Model struct {
	db db
}

// New - return new model
func New(db db) *Model {
	return &Model{
		db: db,
	}
}

// Users - return Users array
func (m *Model) Users() ([]*User, error) {
	return m.db.GetAllUsers()
}

// Products - return Products array
func (m *Model) Products() ([]*Product, error) {
	return m.db.GetAllProducts()
}
