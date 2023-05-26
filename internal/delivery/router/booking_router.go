package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
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
		courts.POST("/", b.bookingHandler.CreateBooking)
		courts.PUT("/update", b.bookingHandler.UpdateBookingPaymentStatus)
		courts.PATCH("/cancel", b.bookingHandler.CancelBooking)
		courts.GET("/detail", b.bookingHandler.GetBookingByOrderNumber)
		courts.GET("/court", b.bookingHandler.GetAllBooking)
		courts.GET("/", b.bookingHandler.GetAllBooking)
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
