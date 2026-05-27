package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
)

type GetUserServiceI interface {
	GetUser(id int64) (*model.User, error)
}

type getUserService struct {
	repo repository.GetUserRepoI
}

func NewGetUserService(repo repository.GetUserRepoI) GetUserServiceI {
	return &getUserService{
		repo:repo,
	}
}

func (s *getUserService) GetUser(id int64) (*model.User, error) {
	return s.repo.GetUser(id)
}