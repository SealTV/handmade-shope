package db

import (
	"database/sql"

	"github.com/SealTV/handmade-shope/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitSqlite(cfg Config) (*pgDb, error) {
	if dbConn, err := sqlx.Connect("postgres", cfg.ConnectString); err != nil {
		return nil, err
	} else {
		p := &pgDb{dbConn: dbConn}
		if err := p.dbConn.Ping(); err != nil {
			return nil, err
		}
		if err := p.createTablesIfNotExist(); err != nil {
			return nil, err
		}
		if err := p.prepareSqlStatements(); err != nil {
			return nil, err
		}
		return p, nil
	}
}

type pgDb struct {
	dbConn *sqlx.DB

	sqlSelectPeople *sqlx.Stmt
	sqlInsertPerson *sqlx.NamedStmt
	sqlSelectPerson *sql.Stmt
}

func (p *pgDb) createTablesIfNotExist() error {
	create_sql := `
	 CREATE TABLE IF NOT EXISTS table_name
(
  id INT,
  name VARCHAR NOT NULL,
  description VARCHAR,
  image BINARY(2048),
  price INT DEFAULT 0,
  create_on DATE DEFAULT CURRENT_DATE,
  update_on DATE DEFAULT CURRENT_DATE
);
CREATE UNIQUE INDEX IF NOT EXISTS table_name_id_uindex ON table_name (id);
CREATE UNIQUE INDEX IF NOT EXISTS table_name_name_uindex ON table_name (name);

CREATE TABLE IF NOT EXISTS users
(
    id integer PRIMARY KEY AUTOINCREMENT NOT NULL,
    login varchar NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL
);
    `
	if rows, err := p.dbConn.Query(create_sql); err != nil {
		return err
	} else {
		rows.Close()
	}
	return nil
}

func (p *pgDb) prepareSqlStatements() (err error) {

	if p.sqlSelectPeople, err = p.dbConn.Preparex(
		"SELECT id, first, last FROM people",
	); err != nil {
		return err
	}
	if p.sqlInsertPerson, err = p.dbConn.PrepareNamed(
		"INSERT INTO people (first, last) VALUES (:first, :last) " +
			"RETURNING id, first, last",
	); err != nil {
		return err
	}
	if p.sqlSelectPerson, err = p.dbConn.Prepare(
		"SELECT id, first, last FROM people WHERE id = $1",
	); err != nil {
		return err
	}

	return nil
}

func (p *pgDb) SelectPeople() ([]*model.User, error) {
	people := make([]*model.User, 0)
	if err := p.sqlSelectPeople.Select(&people); err != nil {
		return nil, err
	}
	return people, nil
}
