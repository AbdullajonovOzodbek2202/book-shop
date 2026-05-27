package service

import (
	"app/src/main/app/model"
	"app/src/main/app/repository"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginUserServiceI interface {
	LoginUser(req *model.LoginUserModel) (string, error)
}

type loginUserService struct {
	repo repository.LoginUserRepoI
}

func NewLoginUserService(repo repository.LoginUserRepoI) LoginUserServiceI {
	return &loginUserService{repo: repo}
}

func (s *loginUserService) LoginUser(req *model.LoginUserModel) (string, error) {
	user, err := s.repo.LoginUser(req)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "" , fmt.Errorf("Token Yaratishda xatolik!")
	}


	return tokenString, nil
}