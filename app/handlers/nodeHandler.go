package handler

import (
	"github.com/gin-gonic/gin"
	"rest.gtld.test/realTimeApp/app/model"
	"rest.gtld.test/realTimeApp/app/usecases"
	"net/http"
)


type nodeHandler struct {
	NodeUsecaseImp  *usecases.NodeUsecaseImp
}

func NewNodeHandler(nodeUsecase *usecases.NodeUsecaseImp)*nodeHandler{
	return &nodeHandler{
		NodeUsecaseImp: nodeUsecase,
	}
}

func (n *nodeHandler) HnadleLogin(c *gin.Context, ){
	var json model.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if ok := n.NodeUsecaseImp.AuthenticateNode(c, &json); ok {
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
	}
}