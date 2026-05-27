package repository

import (
	"app/src/main/app/model"
	"database/sql"
)

type CreateBookRepoI interface {
	CreateBook(req *model.CreateBookModel) (*model.Book, error)
}

type createBookRepo struct {
	db *sql.DB
}

func CreateBookRepo(db *sql.DB) CreateBookRepoI {
	return &createBookRepo{
		db: db,
	}
}

func (cbr *createBookRepo) CreateBook(req *model.CreateBookModel) (*model.Book, error) {

	query := `
		INSERT INTO books(
			title,
			author,
			quantity
		)
		VALUES($1, $2, $3)
		RETURNING id
	`

	var id int64

	err := cbr.db.QueryRow(
		query,
		req.Title,
		req.Author,
		req.Quantity,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &model.Book{
		Id:       id,
		Title:    req.Title,
		Author:   req.Author,
		Quantity: req.Quantity,
	}, nil
}