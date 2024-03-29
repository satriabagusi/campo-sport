package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/pkg/helper"
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
	AdminGetAllUsers(*gin.Context)
	GetAllTopupHistory(*gin.Context)
	GetAllBookingHistory(*gin.Context)
}

func (u *userHandler) Me(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	userId := user.ID

	result, _ := u.userUsecase.FindUserDetailById(userId)

	userResponse := &res.GetUserProfile{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		UserRole:    user.UserRole,
		Detail:      result,
		IsVerified:  user.IsVerified,
	}

	helper.Response(c, http.StatusOK, "OK", userResponse)
}

func (u *userHandler) UpdateMyPassword(c *gin.Context) {

	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	userPassword := &req.UpdatedPassword{
		Id:       user.ID,
		Password: user.Password,
	}

	if err := c.ShouldBindJSON(&userPassword); err != nil {
		helper.Response(c, http.StatusBadRequest, "password is required!", nil)
		return
	}

	userPassword.Password = utility.Encrypt(userPassword.Password)

	_, err := u.userUsecase.UpdatePassword(userPassword)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Handle validation errors
			validationErrors := make(map[string]string)

			for _, e := range validationErrs {
				validationErrors[e.Field()] = e.Tag()
			}

			helper.Response(c, http.StatusBadRequest, "Validation error", validationErrors)
			return
		}

		helper.Response(c, http.StatusInternalServerError, "Failed to update password", nil)
		return
	}

	helper.Response(c, http.StatusOK, "Sucessfully update password", nil)
}

func (h *userHandler) AdminGetAllUsers(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	result, err := h.userUsecase.AdminGetAllUsers()
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Failed to get data", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)

}

func (u *userHandler) GetAllTopupHistory(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	userId := user.ID

	result, err := u.userUsecase.GetAllTopupHistory(userId)
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)
}

func (u *userHandler) GetAllBookingHistory(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	userId := user.ID

	result, err := u.userUsecase.GetAllBookingHistory(userId)
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)
}
