package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
	"github.com/satriabagusi/campo-sport/internal/delivery/middleware"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type UserRouter struct {
	userHandler handler.UserHandler
	publicRoute *gin.RouterGroup
}

func (u *UserRouter) SetupRouter() {
	u.publicRoute.POST("/login/user", u.userHandler.Login)
	u.publicRoute.POST("/register", u.userHandler.InsertUser)

	users := u.publicRoute.Group("/users")
	{
		users.Use(middleware.Authentication())
		users.POST("/", u.userHandler.InsertUser)
		users.PUT("/updatests", u.userHandler.UpdateUser)
		users.PUT("/updatepw", u.userHandler.UpdatePassword)
		users.DELETE("/:id", u.userHandler.DeleteUser)
		users.GET("/:id", u.userHandler.FindUserById)
		users.GET("/email", u.userHandler.FindUserByEmail)
		users.GET("/username", u.userHandler.FindUserByUsername)
		users.GET("/", u.userHandler.GetAllUsers)
	}

}

func NewUserRouter(publicRoute *gin.RouterGroup, userUsecase usecase.UserUsecase) {
	userHandler := handler.NewUserHandler(userUsecase)
	rt := UserRouter{
		userHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
