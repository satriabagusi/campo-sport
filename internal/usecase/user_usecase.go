package usecase

import (
	"fmt"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type UserUsecase interface {
	InsertUser(*req.User) (*res.User, error)
	FindUserByUsername(string) (*entity.User, error)
	Login(*entity.User) (*entity.User, error)
}

type userUsecase struct {
	userRepository       repository.UserRepository
	userDetailRepository repository.UserDetailRepository
}

func NewUserUsecase(userRepository repository.UserRepository, userDetailRepository repository.UserDetailRepository) UserUsecase {
	return &userUsecase{userRepository,
		userDetailRepository,
	}
}

func (u *userUsecase) InsertUser(user *req.User) (*res.User, error) {
	// minPasswordLenght := utility.GetEnv("MIN_PASSWORD_LENGHT")
	// intMintPasswordLenght, err := strconv.Atoi(minPasswordLenght)
	// if err != nil {
	// 	return nil, err
	// }
	if len(user.Password) < 6 {
		return nil, fmt.Errorf("password must be atleast %d characters", 6)
	}
	//return u.userRepository.InsertUser(user)
	_, err := u.userRepository.InsertUser(user)
	if err != nil {
		return nil, err
	}

	userDetail := &req.UserDetail{
		UserId:  user.Id,
		Balance: 0,
	}
	err = u.userDetailRepository.CreateUserDetail(userDetail)
	if err != nil {
		return nil, err
	}
	return &res.User{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (u *userUsecase) FindUserByUsername(username string) (*entity.User, error) {
	return u.userRepository.FindUserByUsername(username)
}

func (u *userUsecase) Login(user *entity.User) (*entity.User, error) {
	return u.userRepository.FindUserByUsername(user.Username)
}
