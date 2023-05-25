package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	FindUserById(*gin.Context)
	FindUserByEmail(*gin.Context)
	GetAllVoucher(*gin.Context)
	InsertUser(*gin.Context)
	FindUserByUsername(*gin.Context)
	Login(*gin.Context)

	UpdatePassword(*gin.Context)
}
