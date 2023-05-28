package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/pkg/token"
	"github.com/satriabagusi/campo-sport/pkg/utility"
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
	UpdateMyPassword(*gin.Context)
}

func (u *userHandler) Me(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)
	
	

	userResponse := &res.GetAllUser{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		UserRole:    user.UserRole,
	}
	webResponse := res.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (u *userHandler) UpdateMyPassword(c *gin.Context) {

	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	userPassword := &req.UpdatedPassword{
		Id:       user.ID,
		Password: user.Password,
	}

	if err := c.ShouldBindJSON(&userPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userPassword.Password = utility.Encrypt(userPassword.Password)

	pw, err := u.userUsecase.UpdatePassword(userPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update voucher"})
		return
	}

	webResponse := res.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   pw,
	}

	c.JSON(http.StatusOK, webResponse)
}
