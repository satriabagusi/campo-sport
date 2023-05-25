/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_top_up_usecase.go (c) 2023
Desc: description
Created:  2023-05-24T17:38:08.976Z
Modified: !date!
*/

package usecase

import (
	"errors"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type UserTopUpUsecase interface {
	TopUpBalance(newTopUp *entity.UserTopUp) (*entity.UserTopUp, error)
}

type userTopUpUsecase struct {
	userTopUpRepo repository.UserTopUpRepository
}

func (u *userTopUpUsecase) TopUpBalance(newTopUp *entity.UserTopUp) (*entity.UserTopUp, error) {
	if newTopUp.Amount == 0 {
		return newTopUp, errors.New("ammount to top up is 0. top up canceled")
	}

	return u.userTopUpRepo.TopUpBalance(newTopUp)
}

func NewUserTopUpUsecase(userTopUp repository.UserTopUpRepository) UserTopUpUsecase {
	return &userTopUpUsecase{userTopUp}
}
