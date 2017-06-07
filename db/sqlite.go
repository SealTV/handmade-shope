package db

import (
	"database/sql"

	"github.com/SealTV/handmade-shope/model"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDb struct {
	dbConn *sql.DB

	sqlSelectAllUsers *sql.Stmt
	sqlSelectUser     *sql.Stmt
	sqlInsertUser     *sql.Stmt
	sqlUpdateUser     *sql.Stmt
	sqlDeleteUser     *sql.Stmt

	sqlSelectAllProducts *sql.Stmt
	sqlSelectProduct     *sql.Stmt
	sqlInsertProduct     *sql.Stmt
	sqlUpdateProduct     *sql.Stmt
	sqlDeleteProduct     *sql.Stmt
}

// InitSqlite - init sqlite
func InitSqlite(cfg string) (*SqliteDb, error) {
	db, err := sql.Open("sqlite3", cfg)
	if err != nil {
		return nil, err
	}

	result := &SqliteDb{dbConn: db}
	result.createTablesIfNotExist()
	result.prepareSQLStatements()
	return result, nil
}

func (db *SqliteDb) createTablesIfNotExist() error {
	creatSQL := `
	 CREATE TABLE IF NOT EXISTS products 
	 ( 
		 id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, 
		 name INTEGER NOT NULL UNIQUE, 
		 description INTEGER NOT NULL, 
		 image INTEGER, 
		 price INTEGER DEFAULT 0, 
		 create_on INTEGER NOT NULL, 
		 update_on INTEGER NOT NULL
	 );

	CREATE TABLE IF NOT EXISTS users 
	( 
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, 
		login varchar NOT NULL UNIQUE, 
		email INTEGER NOT NULL UNIQUE, 
		password INTEGER NOT NULL 
	);
    `
	_, err := db.dbConn.Exec(creatSQL)
	return err
}

func (db *SqliteDb) prepareSQLStatements() (err error) {

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
