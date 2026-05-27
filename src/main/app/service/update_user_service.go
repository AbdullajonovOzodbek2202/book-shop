package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
)

type UpdateUserServiceI interface {
	UpdateUser(id int64, req *model.UpdateUserModel) error
}

type updateUserService struct {
	repo repository.UpdateUserRepoI
}

func NewUpdateUserService(repo repository.UpdateUserRepoI) UpdateUserServiceI {
	return &updateUserService{
		repo: repo,
	}
}

func (s *updateUserService) UpdateUser(id int64, req *model.UpdateUserModel) error {
	return s.repo.UpdateUser(id, req)
}