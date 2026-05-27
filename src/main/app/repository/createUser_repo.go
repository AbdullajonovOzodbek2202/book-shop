package repository

import (
	"app/src/main/app/model"
	"database/sql"
)

type CreateUserRepoI interface {
    CreateUser(req *model.CreateUserRequest) (*model.User, error)
}


type createUserRepo struct {
	db *sql.DB
}

func CreateUserRepo(db *sql.DB) CreateUserRepoI {
	return &createUserRepo{db: db}
}

func (cur *createUserRepo) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
    query := "INSERT INTO users(name, username, password, gmail) VALUES($1, $2, $3, $4) RETURNING id"

    var id int64
    err := cur.db.QueryRow(query, req.Name, req.UserName, req.Password, req.Gmail).Scan(&id)
    if err != nil {
        return nil, err
    }

    return &model.User{
        Id:       &id,
        Name:     &req.Name,
        UserName: &req.UserName,
        Password: &req.Password,
        Gmail:    &req.Gmail,
    }, nil
}