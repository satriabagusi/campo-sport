package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
	"github.com/satriabagusi/campo-sport/internal/delivery/middleware"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type UserDetailRouter struct {
	userDetailHandler handler.UserDetailHandler
	publicRoute       *gin.RouterGroup
}

func (u *UserDetailRouter) SetupRouter() {
	userDetail := u.publicRoute.Group("/userdetail")
	{
		userDetail.Use(middleware.Authentication())
		userDetail.POST("/", u.userDetailHandler.GetAllUserDetail)
		userDetail.PUT("/:id", u.userDetailHandler.GetAllUserDetail)
		userDetail.GET("/court", u.userDetailHandler.GetAllUserDetail)
	}

}

func NewUserDetailRouter(publicRoute *gin.RouterGroup, userDetailUsecase usecase.UserDetailUsecase) {
	userDetailHandler := handler.NewUserDetailHandler(userDetailUsecase)
	rt := UserDetailRouter{
		userDetailHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
