package repository

import (
	"app/src/main/app/model"
	"database/sql"
)

type GetBookRepoI interface {
	GetBook(id int64) (*model.Book, error)
}

type getBookRepo struct {
	db *sql.DB
}

func NewGetBookrepo(db *sql.DB) GetBookRepoI {
	return &getBookRepo{db: db}
}

func (gb *getBookRepo) GetBook(id int64) (*model.Book, error) {
	query := "SELECT *FROM books WHERE id = $1"

	var book model.Book

	err := gb.db.QueryRow(query, id).Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return &book, nil
}
