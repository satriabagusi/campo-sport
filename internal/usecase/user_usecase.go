package usecase

import (
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type UserUsecase interface {
	FindUserByUsername(string) (*entity.User, error)
	Login(*entity.User) (*entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository}
}

func (u *userUsecase) FindUserByUsername(username string) (*entity.User, error) {
	return u.userRepository.FindUserByUsername(username)
}

func (u *userUsecase) Login(user *entity.User) (*entity.User, error) {
	return u.userRepository.FindUserByUsername(user.Username)
}
