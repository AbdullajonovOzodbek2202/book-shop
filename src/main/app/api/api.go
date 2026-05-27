package api

import (
	"app/src/main/app/api/handler"
	"app/src/main/app/middleware"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Api(router *http.ServeMux, handler *handler.Handler) {

	router.HandleFunc("GET /users", middleware.AuthMiddleware(handler.GetUsers))
	router.HandleFunc("POST /createUser", handler.CreateUser)
	router.HandleFunc("POST /login", handler.LoginUser)
	router.HandleFunc("PUT /users/{id}", handler.UpdateUser)
	router.HandleFunc("DELETE /users/{id}", handler.DeleteUser)
	router.HandleFunc("GET /users/{id}", handler.GetUser)
	router.HandleFunc("GET /books", handler.GetBooks)
	router.HandleFunc("POST /books",handler.CreateBook)
	router.HandleFunc("GET /books/{id}",handler.GetBookById)
	router.HandleFunc("PUT /books/{id}",handler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}",handler.DeleteBook)
	router.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
