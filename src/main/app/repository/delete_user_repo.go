package repository

import "database/sql"

type DeleteUserRepoI interface {
	DeleteUser(id int64) error
}

type deleteUserRepo struct {
	db *sql.DB
}

func NewDeleteUserRepo(db *sql.DB) DeleteUserRepoI {
	return &deleteUserRepo{
		db: db,
	}
}

func (r *deleteUserRepo) DeleteUser(id int64) error {

	query := `
		DELETE FROM users
		WHERE id=$1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}