package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"rest.gtld.test/realTimeApp/app/model"
	"rest.gtld.test/realTimeApp/app/usecases"
)


type userHandler struct {
	userUsecaseImp *usecases.UserUsecaseImp
}

func NewUserHanlder(userUsecase * usecases.UserUsecaseImp) *userHandler {
	return &userHandler{
		userUsecaseImp: userUsecase,
	}
}

func (u userHandler) HandleLogin(c *gin.Context) {
	var json model.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if ok := u.userUsecaseImp.AuthenticateUser(c, &json); ok {
	
		generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username":  json.Username,
			"exp": time.Now().Add(time.Hour * 2).Unix(),
		})
	
		token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
	}
}