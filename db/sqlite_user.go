package db

import "github.com/SealTV/handmade-shope/model"

func (db *SqliteDb) prepareUesrsSQLStatements() (err error) {

	db.sqlSelectAllUsers, err = db.dbConn.Prepare("SELECT id, login, email, password FROM users")
	if err != nil {
		return err
	}

	db.sqlSelectUser, err = db.dbConn.Prepare(`
	SELECT id, login, email, password
	FROM users
	WHERE login = ? AND password = ?
	`)
	if err != nil {
		return err
	}

	db.sqlInsertUser, err = db.dbConn.Prepare(`
	INSERT INTO users(login, email, password) VALUES (?, ?, ?)
	`)
	if err != nil {
		return err
	}

	db.sqlUpdateUser, err = db.dbConn.Prepare(`
	UPDATE users
	SET login = ?, email = ?, password = ?
	WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	db.sqlDeleteUser, err = db.dbConn.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}

	return nil
}

// GetAllUsers - return all users
func (db *SqliteDb) GetAllUsers() ([]*model.User, error) {
	users := make([]*model.User, 0)
	rows, err := db.sqlSelectAllUsers.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (db *SqliteDb) GetUser(login, password string) (*model.User, error) {
	var user model.User
	err := db.sqlSelectUser.QueryRow(&login, &password).
		Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *SqliteDb) SetUser(user *model.User) error {
	res, err := db.sqlInsertUser.Exec(&user.UserName, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = lastID
	return nil
}

func (db *SqliteDb) UpdateUser(user *model.User) error {
	_, err := db.sqlUpdateUser.Exec(&user.UserName, &user.Email, &user.Password, &user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (db *SqliteDb) DeleteUser(user *model.User) error {
	_, err := db.sqlDeleteUser.Exec(&user.Id)
	if err != nil {
		return err
	}

	return nil
}
