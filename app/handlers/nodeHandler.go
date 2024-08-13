package handler

import (
	"net/http"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"rest.gtld.test/realTimeApp/app/model"
	"rest.gtld.test/realTimeApp/app/usecases"
)


type nodeHandler struct {
	nodeUsecaseImp  *usecases.NodeUsecaseImp
}

func NewNodeHandler(nodeUsecase *usecases.NodeUsecaseImp)*nodeHandler{
	return &nodeHandler{
		nodeUsecaseImp: nodeUsecase,
	}
}

func (n *nodeHandler) HandleLogin(c *gin.Context){
	var json model.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if ok := n.nodeUsecaseImp.AuthenticateNode(c, &json); ok {
	
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

func (n nodeHandler) GetCurrenct(username string,user *model.Login){
	n.nodeUsecaseImp.GetLoginNode(username, user)
}

func (n nodeHandler) NodeListHandler(c *gin.Context) {
	allNode := n.nodeUsecaseImp.GetAllNode()

	c.JSON(
		http.StatusOK,
		gin.H{
			"nodes": allNode,
		},
	)
}