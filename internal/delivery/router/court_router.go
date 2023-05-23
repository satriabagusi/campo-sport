package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
	"github.com/satriabagusi/campo-sport/internal/delivery/middleware"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type CourtRouter struct {
	courHandler handler.CourtHandler
	publicRoute *gin.RouterGroup
}

func (u *CourtRouter) SetupRouter() {
	courts := u.publicRoute.Group("/courts")
	{
		courts.Use(middleware.Authentication())
		courts.POST("/", u.courHandler.GetAllCourts)
		courts.PUT("/:id", u.courHandler.GetAllCourts)
		courts.DELETE("/:id", u.courHandler.GetAllCourts)
		courts.GET("/:id", u.courHandler.GetAllCourts)
		courts.GET("/court", u.courHandler.GetAllCourts)
		courts.GET("/", u.courHandler.GetAllCourts)
	}

}

func NewCourtRouter(publicRoute *gin.RouterGroup, courtUsecase usecase.CourtUsecase) {
	courtHandler := handler.NewCourtHandler(courtUsecase)
	rt := CourtRouter{
		courtHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
