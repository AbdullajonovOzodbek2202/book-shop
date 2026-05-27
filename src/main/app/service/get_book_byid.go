package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
)

type GetBookServiceI interface {
	GetBook(id int64) (*model.Book, error)
}

type getBookService struct {
	repo repository.GetBookRepoI
}

func NewGetBookService(repo repository.GetBookRepoI) GetBookServiceI {
	return &getBookService{repo:repo}
}

func (s *getBookService) GetBook(id int64) (*model.Book, error) {
	return s.repo.GetBook(id)
}