package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type UserRouter struct {
	userHandler handler.UserHandler
	publicRoute *gin.RouterGroup
}

func (u *UserRouter) SetupRouter() {
	u.publicRoute.POST("/login/user", u.userHandler.Login)
}

func NewUserRouter(publicRoute *gin.RouterGroup, userUsecase usecase.UserUsecase) {
	userHandler := handler.NewUserHandler(userUsecase)
	rt := UserRouter{
		userHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
