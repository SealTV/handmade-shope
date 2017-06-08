package db

import (
	"database/sql"

	"github.com/SealTV/handmade-shope/model"
	_ "github.com/mattn/go-sqlite3"
	"time"
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

func (db *SqliteDb) prepareProductsSQLStatements() (err error) {

	db.sqlSelectAllProducts, err = db.dbConn.Prepare(`
	SELECT id, name, description, image, price, create_on, update_on
	FROM products`)
	if err != nil {
		return err
	}

	db.sqlSelectProduct, err = db.dbConn.Prepare(`
	SELECT id, name, description, image, price, create_on, update_on
	FROM products 
	WHERE name = ? 
	`)
	if err != nil {
		return err
	}

	db.sqlInsertProduct, err = db.dbConn.Prepare(`
	INSERT INTO products(name, description, image, price) VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}

	db.sqlUpdateProduct, err = db.dbConn.Prepare(`
	UPDATE products
	SET name = ?, description = ?, image = ?, price = ?, update_on = ?
	WHERE id = ?;
	`)
	if err != nil {
		return err
	}

	db.sqlDeleteProduct, err = db.dbConn.Prepare("DELETE FROM products WHERE id = ?")
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
	_, err := db.sqlDeleteUser.Exec(&user.Id )
	if err != nil {
		return err
	}

	return nil
}

func (db *SqliteDb) GetAllProducts() ([]*model.Product, error) {
	products := make([]*model.Product, 0)
	rows, err := db.sqlSelectAllProducts.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Image,
			&product.Price,
			&product.CreatedOn,
			&product.UpdatedOn
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (db *SqliteDb) GetProduct(productName string) (*model.Product, error) {
	var product model.Product
	err := db.sqlSelectProduct.QueryRow(&productName).Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Image,
			&product.Price,
			&product.CreatedOn,
			&product.UpdatedOn
		)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (db *SqliteDb) SetProduct(p *model.Product) error {
	res, err := db.sqlInsertProduct.Exec(
		&p.Name,
		&p.Description,
		&p.Image,
		&p.Price
	)
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

func (db *SqliteDb) UpdateProduct(p *model.Product) error {
	_, err := db.sqlUpdateProduct.Exec(
		&p.Name,
		&p.Description,
		&p.Image,
		&p.Price,
		&time.Now(),
		&p.Id
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *SqliteDb) DeleteProduct(p *model.Product) error {
	_, err := db.sqlDeleteProduct.Exec(&p.Id )
	if err != nil {
		return err
	}

	return nil
}