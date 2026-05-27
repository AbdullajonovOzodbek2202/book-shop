package db

import "database/sql"

const (
	DSN = "postgres://postgres:Ao1234@localhost/bookshop?sslmode=disable"
)

func NewDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}
