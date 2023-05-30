/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_top_up_router.go (c) 2023
Desc: description
Created:  2023-05-24T18:17:47.073Z
Modified: !date!
*/

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
	"github.com/satriabagusi/campo-sport/internal/delivery/middleware"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type UserTopUpRouter struct {
	userTopUpHandler handler.UserTopUpHandler
	publicRoute      *gin.RouterGroup
}

func (r *UserTopUpRouter) SetupRouter() {
	r.publicRoute.POST("balance", r.userTopUpHandler.TopUpBalance)
	r.publicRoute.GET("check", r.userTopUpHandler.CheckBalance)
	// r.publicRoute.POST("withdraw", r.userTopUpHandler.WithdrawBalance)

	usersTopUpBalance := r.publicRoute.Group("/user/balance")
	{
		usersTopUpBalance.Use(middleware.Auth())
		usersTopUpBalance.POST("/top-up", r.userTopUpHandler.TopUpBalance)
		usersTopUpBalance.GET("/check", r.userTopUpHandler.CheckBalance)
		usersTopUpBalance.POST("/withdraw", r.userTopUpHandler.WithdrawBalance)
	}
}

func NewUserTopUpRouter(publicRoute *gin.RouterGroup, userTopUp usecase.UserTopUpUsecase) {
	userTopUpHandler := handler.NewUserTopUpHandler(userTopUp)
	rt := UserTopUpRouter{
		userTopUpHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
