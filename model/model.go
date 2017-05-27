package model

// Model - base type
type Model struct {
	db
}

// New - return new model
func New(db db) *Model {
	return &Model{
		db: db,
	}
}

// Users - return Users array
func (m *Model) Users() ([]*User, error) {
	return m.GetAllUsers()
}
