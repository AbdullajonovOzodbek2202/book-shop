package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
)

type BooksServiceI interface {
	GetBooks() ([]*model.Book, error)
}

type BookService struct {
	bookRepo repository.BookRepoI
}

func NewBookService(bookRepo repository.BookRepoI) BooksServiceI {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (bs *BookService) GetBooks() ([]*model.Book, error) {

	books, err := bs.bookRepo.GetBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}
