package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	InsertUser(*gin.Context)
	FindUserById(*gin.Context)
	FindUserByUsername(*gin.Context)
	FindUserByEmail(*gin.Context)
	GetAllUsers(*gin.Context)
	UpdateUser(*gin.Context)
	UpdatePassword(*gin.Context)
	DeleteUser(*gin.Context)
	Login(*gin.Context)
}
