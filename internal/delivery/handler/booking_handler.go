package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/usecase"
	"github.com/satriabagusi/campo-sport/pkg/token"
)

type BookingHandler interface {
	GetAllBooking(*gin.Context)
	GetBookingByOrderNumber(*gin.Context)
	CreateBooking(*gin.Context)
	UpdateBookingPaymentStatus(*gin.Context)
	CancelBooking(*gin.Context)
}

type bookingHandler struct {
	bookingUsecase usecase.BookingUsecase
}

func NewBookingHandler(bookingUsecase usecase.BookingUsecase) BookingHandler {
	return &bookingHandler{bookingUsecase}
}

func (h *bookingHandler) GetAllBooking(c *gin.Context) {
	res, err := h.bookingUsecase.GetAllBooking()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Get Data Successfully",
		"data":    res,
	})
}

func (h *bookingHandler) CreateBooking(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)
	userId := user.ID

	var booking entity.Booking
	booking.User.Id = userId

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    booking,
		})
		return
	}

	result, err := h.bookingUsecase.CreateBooking(&booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    booking,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"code":    http.StatusAccepted,
		"message": "Booking Created.",
		"data":    result,
	})
}

func (h *bookingHandler) GetBookingByOrderNumber(c *gin.Context) {
	bookingNumber := c.Query("booking_number")

	result, err := h.bookingUsecase.GetBookingByOrderNumber(bookingNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Booking Detail Found",
		"data":    result,
	})
}

func (h *bookingHandler) UpdateBookingPaymentStatus(c *gin.Context) {
	bookingNumber := c.Query("booking_number")

	result, err := h.bookingUsecase.UpdateBookingPaymentStatus(bookingNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Get Payment Status Successfully",
		"data":    result,
	})
}

func (h *bookingHandler) CancelBooking(c *gin.Context) {
	bookingNumber := c.Query("booking_number")

	result, err := h.bookingUsecase.CancelBooking(bookingNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Booking Has Been Cancelled",
		"data":    result,
	})
}
