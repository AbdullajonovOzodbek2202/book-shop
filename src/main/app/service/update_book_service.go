package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
)

type UpdateBookServiceI interface {
	UpdateBook(id int64, req *model.UpdateBookModel) error
}

type updateBookService struct {
	repo repository.UpdateBookRepoI
}

func NewUpdateBookService(repo repository.UpdateBookRepoI) UpdateBookServiceI {
	return &updateBookService{
		repo: repo,
	}
}

func (s *updateBookService) UpdateBook(id int64, req *model.UpdateBookModel) error {
	return s.repo.UpdateBook(id, req)
}
