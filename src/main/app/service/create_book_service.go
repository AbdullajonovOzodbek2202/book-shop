package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
)

type CreateBookServiceI interface {
	CreateBook(req *model.CreateBookModel) (*model.Book, error)
}

type createBookService struct {
	repo repository.CreateBookRepoI
}

func NewCreateBookService(repo repository.CreateBookRepoI) CreateBookServiceI {
	return &createBookService{repo:repo}
}

func (s *createBookService) CreateBook(req *model.CreateBookModel) (*model.Book, error) {
	return s.repo.CreateBook(req)
}