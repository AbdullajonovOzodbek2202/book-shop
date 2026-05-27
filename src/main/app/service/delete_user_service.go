package service

import "app/src/main/app/repository"

type DeleteUserServiceI interface {
	DeleteUser(id int64) error
}

type deleteUserService struct {
	repo repository.DeleteUserRepoI
}

func NewDeleteUserService(repo repository.DeleteUserRepoI) DeleteUserServiceI {
	return &deleteUserService{
		repo: repo,
	}
}

func (s *deleteUserService) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}