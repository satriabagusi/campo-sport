package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/usecase"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInDb, _ := u.userUsecase.FindUserByUsername(user.Username)
	if userInDb != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	user.Password = utility.Encrypt(user.Password)

	result, err := u.userUsecase.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": result})
}
func (u *userHandler) UpdateUser(c *gin.Context) {
	var updateUserSts req.UpdatedStatusUser
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	updateUserSts.Id = idInt

	userInDb, err := u.userUsecase.FindUserById(idInt)
	if userInDb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&updateUserSts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := u.userUsecase.UpdateUserStatus(&updateUserSts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update voucher"})
		return
	}

	webResponse := res.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (u *userHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = u.userUsecase.FindUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "voucher not found"})
		return
	}

	err = u.userUsecase.DeleteUser(&entity.User{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "voucher has been deleted",
	}

	c.JSON(http.StatusOK, webResponse)
}
func (u *userHandler) FindUserById(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := u.userUsecase.FindUserById(id)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (u *userHandler) FindUserByEmail(c *gin.Context) {
	emailUser := c.Query("email")

	result, err := u.userUsecase.FindUserByEmail(emailUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	webResponse := res.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (u *userHandler) GetAllUsers(c *gin.Context) {
	result, err := u.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := res.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (u *userHandler) FindUserByUsername(c *gin.Context) {
	username := c.Query("username")

	result, err := u.userUsecase.FindUserByUsername(username)
	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kesalahan server"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *userHandler) Login(c *gin.Context) {
	var login req.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user res.GetUserByUsername
	userInDb, err := u.userUsecase.FindUserByUsernameLogin(login.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	//memverifikasi apakah password yang dimasukkan sama di database dengan helper VerifyPassword
	if err := utility.VerifyPassword(userInDb.Password, login.Password); err != nil {
		log.Println(user.Password, login.Password)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := token.CreateToken(userInDb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}

func (u *userHandler) UpdatePassword(c *gin.Context) {

	var userUpdatePassword req.UpdatedPassword
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	userUpdatePassword.Id = idInt

	userInDb, _ := u.userUsecase.FindUserById(idInt)
	if userInDb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&userUpdatePassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userUpdatePassword.Password = utility.Encrypt(userUpdatePassword.Password)

	_, err := u.userUsecase.UpdatePassword(&userUpdatePassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update voucher"})
		return
	}

	webResponse := res.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Password sucessfully updated",
	}

	c.JSON(http.StatusOK, webResponse)
}

// func (u *userHandler) Login(c *gin.Context) {
// 	var user entity.User

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	userIndDb, err := u.userUsecase.FindUserByUsername(user.Username)

// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if user.Username != userIndDb.Username && user.Password != userIndDb.Password {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	secretKey := utility.GetEnv("SECRET_KEY")
// 	expireTmeInt, err := strconv.Atoi(utility.GetEnv("TOKEN_EXPIRE_TIME_IN_MINUTES"))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	expireAt := time.Now().Add(time.Minute * time.Duration(expireTmeInt))

// 	tokenString, err := jwt.GenerateToken(int64(user.Id), expireAt, []byte(secretKey))

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"token": tokenString})

// }
