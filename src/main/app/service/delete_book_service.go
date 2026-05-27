package service

import "app/src/main/app/repository"

type DeleteBookServiceI interface {
	DeleteBook(id int64) error
}

type deleteBookService struct {
	repo repository.DeleteBookRepoI
}

func NewDeleteBookService(repo repository.DeleteBookRepoI) DeleteBookServiceI {
	return &deleteBookService{repo: repo}
}

func (s *deleteBookService) DeleteBook(id int64) error {
	return s.repo.DeleteBook(id)
}
