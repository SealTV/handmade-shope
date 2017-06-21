package db

import (
	"database/sql"
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
	result.prepareUesrsSQLStatements()
	result.prepareProductsSQLStatements()
	return result, nil
}

func (db *SqliteDb) Close() {
	db.dbConn.Close()
}

func (db *SqliteDb) createTablesIfNotExist() error {
	creatSQL := `
	 CREATE TABLE IF NOT EXISTS products 
	 ( 
		 id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, 
		 name varchar NOT NULL UNIQUE, 
		 description varchar NOT NULL, 
		 image varchar, 
		 price INTEGER DEFAULT 0, 
		 create_on DATETIME NOT NULL, 
		 update_on DATETIME NOT NULL
	 );

	CREATE TABLE IF NOT EXISTS users 
	( 
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, 
		login varchar NOT NULL UNIQUE, 
		email varchar NOT NULL UNIQUE, 
		password varchar NOT NULL 
	);
    `
	_, err := db.dbConn.Exec(creatSQL)
	return err
}
