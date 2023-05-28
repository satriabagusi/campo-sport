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
		courts.Use(middleware.Auth())
		courts.GET("/:id", u.courHandler.FindCourtByID)
		courts.GET("/search", u.courHandler.FindCourtByCourtName)
		courts.GET("/", u.courHandler.GetAllCourts)
	}

	admin := u.publicRoute.Group("/admin/courts")
	{
		courts.Use(middleware.Auth())
		admin.POST("/", u.courHandler.InsertCourt)
		admin.PUT("/edit", u.courHandler.UpdateCourt)
		admin.DELETE("/:id", u.courHandler.DeleteCourt)
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
