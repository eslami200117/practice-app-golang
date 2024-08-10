package usecases

import (
	"github.com/gin-gonic/gin"
)

type NodeUsecase interface {
	AuthenticateNode(c *gin.Context, username string, password string) bool
}
