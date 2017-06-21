package db

import (
	"github.com/SealTV/handmade-shope/model"
	time "time"
)

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
			&product.UpdatedOn,
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
		&product.UpdatedOn,
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
		&p.Price,
	)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	p.Id = lastID
	return nil
}

func (db *SqliteDb) UpdateProduct(p *model.Product) error {
	time := time.Now()
	_, err := db.sqlUpdateProduct.Exec(
		&p.Name,
		&p.Description,
		&p.Image,
		&p.Price,
		&time,
		&p.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *SqliteDb) DeleteProduct(p *model.Product) error {
	_, err := db.sqlDeleteProduct.Exec(&p.Id)
	if err != nil {
		return err
	}

	return nil
}
