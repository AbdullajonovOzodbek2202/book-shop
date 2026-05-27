package repository

import "database/sql"

type DeleteBookRepoI interface {
	DeleteBook(id int64) (error)
}

type deleteBookRepo struct {
	db *sql.DB
}

func NewDeleteBookRepo (db *sql.DB) DeleteBookRepoI {
	return &deleteBookRepo{db : db}
}

func (r *deleteBookRepo) DeleteBook (id int64) error {

	query := "DELETE FROM books WHERE id = $1;"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
