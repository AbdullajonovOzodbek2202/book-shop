package repository

import (
	"app/src/main/app/model"
	"database/sql"
)

type UpdateUserRepoI interface {
	UpdateUser(id int64, req *model.UpdateUserModel) error
}

type updateUserRepo struct {
	db *sql.DB
}

func UpdateUserRepo(db *sql.DB) UpdateUserRepoI {
	return &updateUserRepo{
		db: db,
	}
}

func (ur *updateUserRepo) UpdateUser(id int64, req *model.UpdateUserModel) error {

	query := `
		UPDATE users
		SET
			name=$1,
			username=$2,
			password=$3,
			gmail=$4
		WHERE id=$5
	`

	_, err := ur.db.Exec(
		query,
		req.Name,
		req.UserName,
		req.Password,
		req.Gmail,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}