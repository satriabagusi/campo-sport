package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type CourtHandler interface {
	GetAllCourts(*gin.Context)
}

type courtHandler struct {
	courtUsecase usecase.CourtUsecase
}

func NewCourtHandler(courtUsecase usecase.CourtUsecase) CourtHandler {
	return &courtHandler{courtUsecase}
}

func (h *courtHandler) GetAllCourts(c *gin.Context) {

}
