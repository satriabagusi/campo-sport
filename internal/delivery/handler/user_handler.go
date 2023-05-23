package handler

import (
	"net/http"
	"strconv"
	"time"

	jwt "github.com/eulbyvan/auth-go"
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/usecase"
	"github.com/satriabagusi/campo-sport/pkg/utility"
)

type UserHandler interface {
	FindUserByUsername(*gin.Context)
	Login(*gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase}
}

func (u *userHandler) FindUserByUsername(c *gin.Context) {
	username := c.Query("username")

	result, err := u.userUsecase.FindUserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func (u *userHandler) Login(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIndDb, err := u.userUsecase.FindUserByUsername(user.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credential"})
		return
	}

	if user.Username != userIndDb.Username && user.Password != userIndDb.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	secretKey := utility.GetEnv("SECRET_KEY")
	expireTmeInt, err := strconv.Atoi(utility.GetEnv("TOKEN_EXPIRE_TIME_IN_MINUTES"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	expireAt := time.Now().Add(time.Minute * time.Duration(expireTmeInt))

	tokenString, err := jwt.GenerateToken(int64(user.Id), expireAt, []byte(secretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
