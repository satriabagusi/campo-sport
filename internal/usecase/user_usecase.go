package usecase

import (
	"fmt"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type UserUsecase interface {
	UpdateUserStatus(*req.UpdatedStatusUser) (*req.UpdatedStatusUser, error)
	DeleteUser(*entity.User) error
	FindUserById(int) (*res.GetUserByID, error)
	FindUserByEmail(string) (*res.GetUserByUsername, error)
	GetAllUsers() ([]res.GetAllUser, error)

	InsertUser(user *req.User) (*res.User, error)
	FindUserByUsername(string) (*res.GetUserByUsername, error)
	Login(*entity.User) (*res.GetUserByUsername, error)
	UpdatePassword(*req.UpdatedPassword) (*req.UpdatedPassword, error)
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

	if len(user.Password) < 6 {
		return nil, fmt.Errorf("password must be atleast %d characters", 6)
	}

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

func (u *userUsecase) FindUserByUsername(username string) (*res.GetUserByUsername, error) {
	return u.userRepository.FindUserByUsername(username)
}

func (u *userUsecase) Login(user *entity.User) (*res.GetUserByUsername, error) {
	return u.userRepository.FindUserByUsername(user.Username)
}

func (u *userUsecase) UpdateUserStatus(updatedUser *req.UpdatedStatusUser) (*req.UpdatedStatusUser, error) {
	return u.userRepository.UpdateUserStatus(updatedUser)
}

func (u *userUsecase) DeleteUser(deletedUser *entity.User) error {
	return u.userRepository.DeleteUser(deletedUser)
}

func (u *userUsecase) FindUserById(id int) (*res.GetUserByID, error) {
	return u.userRepository.FindUserById(id)
}

func (u *userUsecase) FindUserByEmail(email string) (*res.GetUserByUsername, error) {
	return u.userRepository.FindUserByEmail(email)
}

func (u *userUsecase) GetAllUsers() ([]res.GetAllUser, error) {
	return u.userRepository.GetAllUsers()
}
func (u *userUsecase) UpdatePassword(updatePw *req.UpdatedPassword) (*req.UpdatedPassword, error) {
	return u.userRepository.UpdatePassword(updatePw)
}
