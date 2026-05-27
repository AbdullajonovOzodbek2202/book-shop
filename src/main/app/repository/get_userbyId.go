package repository

import (
	"app/src/main/app/model"
	"database/sql"
)

type GetUserRepoI interface {
	GetUser(id int64) (*model.User, error)
}

type getUserRepo struct {
	db *sql.DB
}

func NewGetUserRepo(db *sql.DB) GetUserRepoI {
	return &getUserRepo{
		db: db,
	}
}

func (gu *getUserRepo) GetUser(id int64) (*model.User, error) {

	query := `
		SELECT
			id,
			name,
			username,
			gmail
		FROM users
		WHERE id=$1
	`

	var user model.User

	err := gu.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Name,
		&user.UserName,
		&user.Gmail,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}