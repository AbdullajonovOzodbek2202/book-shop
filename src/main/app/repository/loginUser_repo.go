package repository

import (
	"app/src/main/app/model"
	"database/sql"
	"fmt"
)

type LoginUserRepoI interface {
	LoginUser(req *model.LoginUserModel) (*model.User, error)
}

type loginUserRepo struct{
	db *sql.DB
}

func LoginUserRepo(db *sql.DB) LoginUserRepoI {
	return &loginUserRepo{db:db}
}

func (lur *loginUserRepo) LoginUser(req *model.LoginUserModel) (*model.User, error) {
	query := "SELECT id, name, username, password, gmail FROM users WHERE username = $1 AND password = $2"

	var user model.User
	err := lur.db.QueryRow(query, req.UserName, req.Password).Scan(
		&user.Id,
		&user.Name,
		&user.UserName,
		&user.Password,
		&user.Gmail,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("username yoki password xato!")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}