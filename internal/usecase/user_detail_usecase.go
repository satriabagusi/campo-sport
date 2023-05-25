package usecase

import (
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type UserDetailUsecase interface {
	GetAllUserDetail() ([]entity.UserDetail, error)
}

type userDetailUsecase struct {
	userDetailRepository repository.UserDetailRepository
}

func NewUserDetailUsecase(userDetailRepository repository.UserDetailRepository) UserDetailUsecase {
	return &userDetailUsecase{userDetailRepository}
}

func (u *userDetailUsecase) GetAllUserDetail() ([]entity.UserDetail, error) {
	return u.userDetailRepository.GetAllUserDetail()
}
