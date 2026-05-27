package handler

import "app/src/main/app/service"

type Handler struct {
	UserService        service.UserServiceI
	CreateUserService  service.CreateUserServiceI
	LoginUserService   service.LoginUserServiceI
	UpdateUserService  service.UpdateUserServiceI
	DeleteUserService  service.DeleteUserServiceI
	GetUserService     service.GetUserServiceI
	GetBookService     service.BooksServiceI
	CreateBookService  service.CreateBookServiceI
	GetBookByIdService service.GetBookServiceI
	UpdateBookService  service.UpdateBookServiceI
	DeleteBookService  service.DeleteBookServiceI
}
