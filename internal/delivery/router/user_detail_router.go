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
		userDetail.Use(middleware.Auth())
		userDetail.POST("/", u.userDetailHandler.GetAllUserDetail)
		userDetail.PUT("/upload/", u.userDetailHandler.UploadCredential)
		userDetail.GET("/court", u.userDetailHandler.GetAllUserDetail)
	}

}

func NewUserDetailRouter(publicRoute *gin.RouterGroup, userDetailUsecase usecase.UserDetailUsecase, userUsecase usecase.UserUsecase) {
	userDetailHandler := handler.NewUserDetailHandler(userDetailUsecase, userUsecase)
	rt := UserDetailRouter{
		userDetailHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
