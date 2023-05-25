/*
Author: Satria Bagus(satria.bagus18@gmail.com)
court_usecase.go (c) 2023
Desc: description
Created:  2023-05-23T08:02:40.682Z
Modified: !date!
*/

package usecase

import (
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type CourtUsecase interface {
	FindCourtById(id int) (*entity.Court, error)
}

type courtUsecase struct {
	courtRepo repository.CourtRepository
}

func (b *courtUsecase) FindCourtById(id int) (*entity.Court, error) {
	return b.courtRepo.FindCourtById(id)
}

func NewCourtUsecase(courtRepo repository.CourtRepository) CourtUsecase {
	return &courtUsecase{courtRepo}
}
