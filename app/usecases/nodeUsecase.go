package usecases

import (
	"github.com/gin-gonic/gin"
	"rest.gtld.test/realTimeApp/app/model"
)

type NodeUsecase interface {
	AuthenticateNode(c *gin.Context, username string, password string) bool
	GetLoginUser(username string, user *model.Login)
}
