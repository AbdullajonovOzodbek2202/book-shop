package repository

import (
	"app/src/main/app/model"
	"database/sql"
)

type UpdateBookRepoI interface {
	UpdateBook(id int64, req *model.UpdateBookModel) error
}

type updateBookRepo struct {
	db *sql.DB
}

func UpdateBookRepo(db *sql.DB) UpdateBookRepoI {
	return &updateBookRepo{
		db: db,
	}
}

func (ub *updateBookRepo) UpdateBook(id int64, req *model.UpdateBookModel) error {

	query := `
		UPDATE books
		SET
			title=$1,
			author=$2,
			quantity=$3
		WHERE id=$4
	`

	_, err := ub.db.Exec(
		query,
		req.Title,
		req.Author,
		req.Quantity,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}