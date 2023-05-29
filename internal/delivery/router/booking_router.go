package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
	"github.com/satriabagusi/campo-sport/internal/delivery/middleware"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type BookingRouter struct {
	bookingHandler handler.BookingHandler
	publicRoute    *gin.RouterGroup
}

func (b *BookingRouter) SetupRouter() {
	courts := b.publicRoute.Group("/booking")
	{
		//courts.Use(middleware.Authentication())
		courts.Use(middleware.Auth())
		courts.POST("/", b.bookingHandler.CreateBooking)
		courts.PUT("/update", b.bookingHandler.UpdateBookingPaymentStatus)
		courts.PATCH("/cancel", b.bookingHandler.CancelBooking)
		courts.GET("/detail", b.bookingHandler.GetBookingByOrderNumber)
	}

}

func NewBookingRouter(publicRoute *gin.RouterGroup, bookingUsecase usecase.BookingUsecase) {
	bookingHandler := handler.NewBookingHandler(bookingUsecase)
	rt := BookingRouter{
		bookingHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
