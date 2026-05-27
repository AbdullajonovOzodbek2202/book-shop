package repository

import (
	"app/src/main/app/model"
	"database/sql"
)

type UserRepoI interface {
	GetUsers() ([]*model.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoI {
	return &userRepo{db: db}
}

func (ur *userRepo) GetUsers() ([]*model.User, error) {

	query := "SELECT id, name, username, password, gmail FROM users;"

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.UserName,
			&user.Password,
			&user.Gmail,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
