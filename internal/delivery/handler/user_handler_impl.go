package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/usecase"
	"github.com/satriabagusi/campo-sport/pkg/helper"
	"github.com/satriabagusi/campo-sport/pkg/token"
	"github.com/satriabagusi/campo-sport/pkg/utility"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase}
}
func (u *userHandler) InsertUser(c *gin.Context) {
	var user req.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request! data required", nil)
		return
	}

	userInDb, _ := u.userUsecase.FindUserByUsername(user.Username)
	if userInDb != nil {
		helper.Response(c, http.StatusConflict, "User already exist", nil)
		return
	}

	user.Password = utility.Encrypt(user.Password)

	result, err := u.userUsecase.InsertUser(&user)
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

		helper.Response(c, http.StatusInternalServerError, "Failed to register!", nil)
		return
	}

	helper.Response(c, http.StatusCreated, "OK", result)
}
func (u *userHandler) UpdateUser(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorizes", nil)
		return
	}

	var updateUserSts req.UpdatedStatusUser
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	updateUserSts.Id = idInt

	userInDb, _ := u.userUsecase.FindUserById(idInt)
	if userInDb == nil {
		helper.Response(c, http.StatusNotFound, "User not found!", nil)
		return
	}

	if err := c.ShouldBindJSON(&updateUserSts); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request ! failed to update user", nil)
		return
	}
	result, err := u.userUsecase.UpdateUserStatus(&updateUserSts)
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Server error", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)

}

func (u *userHandler) DeleteUser(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "server error!", nil)
		return
	}

	_, err = u.userUsecase.FindUserById(id)
	if err != nil {

		helper.Response(c, http.StatusNotFound, "user not found", nil)
		return
	}

	err = u.userUsecase.DeleteUser(&entity.User{Id: id})
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Server error", nil)
		return
	}

	helper.Response(c, http.StatusOK, "user has been deleted", nil)
}
func (u *userHandler) FindUserById(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request!", nil)
		return
	}

	result, err := u.userUsecase.FindUserById(id)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "user doesn't exist", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)
}

func (u *userHandler) FindUserByEmail(c *gin.Context) {
	emailUser := c.Query("email")

	result, err := u.userUsecase.FindUserByEmail(emailUser)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "user not found", nil)
		return
	}
	helper.Response(c, http.StatusOK, "OK", result)
}

func (u *userHandler) GetAllUsers(c *gin.Context) {
	result, err := u.userUsecase.GetAllUsers()
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Failed to get data", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)
}

func (u *userHandler) FindUserByUsername(c *gin.Context) {
	username := c.Query("username")

	result, err := u.userUsecase.FindUserByUsername(username)
	if result == nil {
		helper.Response(c, http.StatusNotFound, "User not found", nil)
		return
	}
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Server error!", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)
}

func (u *userHandler) Login(c *gin.Context) {
	var login req.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request! data required", nil)
		return
	}

	var user res.GetUserByUsername
	userInDb, err := u.userUsecase.FindUserByUsernameLogin(login.Username)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "username or password are wrong", nil)
		return
	}

	//memverifikasi apakah password yang dimasukkan sama di database dengan helper VerifyPassword
	if err := utility.VerifyPassword(userInDb.Password, login.Password); err != nil {
		log.Println(user.Password, login.Password)
		helper.Response(c, http.StatusNotFound, "username or password are wrong", nil)
		return
	}
	//fmt.Println(userInDb.UserRole)
	tokenString, err := token.CreateToken(userInDb)

	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Server error!", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", tokenString)

}

func (u *userHandler) UpdatePassword(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var userUpdatePassword req.UpdatedPassword
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	userUpdatePassword.Id = idInt

	userInDb, _ := u.userUsecase.FindUserById(idInt)
	if userInDb == nil {
		helper.Response(c, http.StatusNotFound, "User not found", nil)
		return
	}

	if err := c.ShouldBindJSON(&userUpdatePassword); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request! data required", nil)
		return
	}

	userUpdatePassword.Password = utility.Encrypt(userUpdatePassword.Password)

	_, err := u.userUsecase.UpdatePassword(&userUpdatePassword)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Handle validation errors
			validationErrors := make(map[string]string)

			for _, e := range validationErrs {
				validationErrors[e.Field()] = e.Tag()
			}

			helper.Response(c, http.StatusBadRequest, "Validation Error", validationErrors)
			return
		}

		helper.Response(c, http.StatusInternalServerError, "Server Error!", nil)
		return
	}

	helper.Response(c, http.StatusOK, "Password sucessfully updated", nil)
}
