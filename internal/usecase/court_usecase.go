package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type CourtUsecase interface {
	InsertCourt(*entity.Court) (*res.Court, error)
	UpdateCourt(*entity.Court) (*res.Court, error)
	DeleteCourt(*entity.Court) error
	FindCourtById(int) (*res.Court, error)
	FindCourtByCourt(string) (*res.Court, error)
	GetAllCourts() ([]entity.Court, error)
}

type courtUsecase struct {
	courtRepository repository.CourtRepository
	Validate        *validator.Validate
}

func NewCourtUsecase(courtRepository repository.CourtRepository, validate *validator.Validate) CourtUsecase {
	return &courtUsecase{courtRepository, validate}
}

func (u *courtUsecase) InsertCourt(newCourt *entity.Court) (*res.Court, error) {
	err := u.Validate.Struct(newCourt)
	if err != nil {
		return nil, err
	}
	return u.courtRepository.InsertCourt(newCourt)
}

func (u *courtUsecase) UpdateCourt(updatedCourt *entity.Court) (*res.Court, error) {
	err := u.Validate.Struct(updatedCourt)
	if err != nil {
		return nil, err
	}
	return u.courtRepository.UpdateCourt(updatedCourt)
}

func (u *courtUsecase) DeleteCourt(deletedCourt *entity.Court) error {
	return u.courtRepository.DeleteCourt(deletedCourt)
}

func (u *courtUsecase) FindCourtById(id int) (*res.Court, error) {
	return u.courtRepository.FindCourtById(id)
}

func (u *courtUsecase) FindCourtByCourt(courtName string) (*res.Court, error) {
	return u.courtRepository.FindCourtByCourt(courtName)
}

func (u *courtUsecase) GetAllCourts() ([]entity.Court, error) {
	return u.courtRepository.GetAllCourts()
}
