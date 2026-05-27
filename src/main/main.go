package main

import (
	"app/src/main/app/api"
	"app/src/main/app/api/handler"
	"app/src/main/app/repository"
	"app/src/main/app/service"
	"app/src/main/dependencies/db"
	"log"
	"net/http"

	_ "app/docs"

	_ "github.com/lib/pq"
)

// @title Bookshop API
// @version 1.0
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db, err := db.NewDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := http.NewServeMux()

	{
		userRepo := repository.NewUserRepo(db)
		userService := service.NewUserService(userRepo)

		createUserRepo := repository.CreateUserRepo(db)
		createUserService := service.NewCreateUserService(createUserRepo)

		loginUserRepo := repository.LoginUserRepo(db)
		loginUserService := service.NewLoginUserService(loginUserRepo)

		updateUserRepo := repository.UpdateUserRepo(db)
		updateUserService := service.NewUpdateUserService(updateUserRepo)

		deleteUserRepo := repository.NewDeleteUserRepo(db)
		deleteUserService := service.NewDeleteUserService(deleteUserRepo)

		getUserRepo := repository.NewGetUserRepo(db)
		getUserService := service.NewGetUserService(getUserRepo)

		getBooks := repository.NewBookRepo(db)
		getBookService := service.NewBookService(getBooks)

		createBookRepo := repository.CreateBookRepo(db)
		createBookService := service.NewCreateBookService(createBookRepo)

		getBookByIdrepo := repository.NewGetBookrepo(db)
		getBookByIdService := service.NewGetBookService(getBookByIdrepo)

		updateBookRepo := repository.UpdateBookRepo(db)
		updateBookService := service.NewUpdateBookService(updateBookRepo)

		deleteBookRepo := repository.NewDeleteBookRepo(db)
		deleteBookService := service.NewDeleteBookService(deleteBookRepo)

		handler := &handler.Handler{
			UserService:        userService,
			CreateUserService:  createUserService,
			LoginUserService:   loginUserService,
			UpdateUserService:  updateUserService,
			DeleteUserService:  deleteUserService,
			GetUserService:     getUserService,
			GetBookService:     getBookService,
			CreateBookService:  createBookService,
			GetBookByIdService: getBookByIdService,
			UpdateBookService:  updateBookService,
			DeleteBookService:  deleteBookService,
		}

		api.Api(router, handler)
	}

	log.Println("Service is running on :8080")
	http.ListenAndServe(":8080", router)

}
