package usecases

import (
	"github.com/gin-gonic/gin"
	"rest.gtld.test/realTimeApp/app/model"
	repository "rest.gtld.test/realTimeApp/app/repositories"
)


type NodeUsecaseImp struct {
	nodeRepo *repository.WeatherPostgresRepo
}


func NewNodeImp(repo * repository.WeatherPostgresRepo) *NodeUsecaseImp{
	return &NodeUsecaseImp{
		nodeRepo: repo,
	}
}

func (nu *NodeUsecaseImp) AuthenticateNode(c *gin.Context, in *model.Login) bool{
	result := nu.nodeRepo.Authen(in)
	// if result{
		// login, create lo
	// }
	return result
}

func (nu *NodeUsecaseImp) GetLoginUser(username string, user *model.Login){
	nu.nodeRepo.GetUser(username, user)
}