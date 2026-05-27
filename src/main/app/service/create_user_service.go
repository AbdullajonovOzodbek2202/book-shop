package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
)

type CreateUserServiceI interface {
	CreateUser(req *model.CreateUserRequest) (*model.User, error)
}

type createUserService struct {
    repo repository.CreateUserRepoI
}

func NewCreateUserService(repo repository.CreateUserRepoI) CreateUserServiceI {
    return &createUserService{repo: repo}
}

func (s *createUserService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
    return s.repo.CreateUser(req)
}
