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
	result := nu.nodeRepo.AuthenNode(in)
	// if result{
		// login, create lo
	// }
	return result
}

func (nu *NodeUsecaseImp) GetLoginNode(username string, user *model.Login){
	nu.nodeRepo.GetNode(username, user)
}

func (nu *NodeUsecaseImp) GetAllNode() [] model.APInode{
	allNode := nu.nodeRepo.GetAllNode()
	var allAPInode []model.APInode

	for _, n := range allNode{
		value := model.APInode{
			Username: n.Username,
			Status: n.Status,
			LastUpdate: timeAgo(n.LastUpdata),
		}
		allAPInode = append(allAPInode, value)
	}
	return allAPInode
}

func (nu *NodeUsecaseImp) AddSource(username string, addSource string, password string) {
	nu.nodeRepo.AddSource(username, addSource, password)
}