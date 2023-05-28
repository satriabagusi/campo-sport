package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/pkg/token"
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
	Me(*gin.Context)
}

func (u *userHandler) Me(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	userResponse := &entity.User{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
	webResponse := res.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}
