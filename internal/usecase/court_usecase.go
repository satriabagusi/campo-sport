package usecase

import (
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type CourtUsecase interface {
	GetAllCourts() ([]entity.Court, error)
}

type courtUsecase struct {
	courtRepository repository.CourtRepository
}

func NewCourtUsecase(courtRepository repository.CourtRepository) CourtUsecase {
	return &courtUsecase{courtRepository}
}

func (u *courtUsecase) GetAllCourts() ([]entity.Court, error) {
	return u.courtRepository.GetAllCourts()
}
