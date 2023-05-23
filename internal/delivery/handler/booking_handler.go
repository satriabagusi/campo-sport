package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type BookingHandler interface {
	GetAllBooking(*gin.Context)
}

type bookingHandler struct {
	bookingUsecase usecase.BookingUsecase
}

func NewBookingHandler(bookingUsecase usecase.BookingUsecase) BookingHandler {
	return &bookingHandler{bookingUsecase}
}

func (h *bookingHandler) GetAllBooking(c *gin.Context) {

}
