package repository

import (
	"app/src/main/app/model"
	"database/sql"
)

type BookRepoI interface {
	GetBooks() ([]*model.Book, error)
}

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) BookRepoI {
	return &bookRepo{
		db: db,
	}
}

func (br *bookRepo) GetBooks() ([]*model.Book, error) {

	query := `
		SELECT
			id,
			title,
			author,
			quantity
		FROM books
	`

	rows, err := br.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*model.Book

	for rows.Next() {
		var book model.Book

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.Quantity,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
