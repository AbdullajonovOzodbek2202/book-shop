package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
)

type UserServiceI interface {
	GetUsers() ([]*model.User, error)
}

type UserService struct {
	userrepo repository.UserRepoI
}

func NewUserService(userrepo repository.UserRepoI) UserServiceI {
	return &UserService{userrepo: userrepo}
}

func (us *UserService) GetUsers() ([]*model.User,error) {
	users, err := us.userrepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}