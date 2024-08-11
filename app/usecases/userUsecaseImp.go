package usecases

import (
	"github.com/gin-gonic/gin"
	"rest.gtld.test/realTimeApp/app/model"
	repository "rest.gtld.test/realTimeApp/app/repositories"
)


type UserUsecaseImp struct {
	repo *repository.WeatherPostgresRepo
}

func NewUserImp(repo *repository.WeatherPostgresRepo) *UserUsecaseImp{
	return &UserUsecaseImp{
		repo: repo,
	}
}

func (u UserUsecaseImp) AuthenticateUser(c *gin.Context, user *model.Login) bool {
	result := u.repo.AuthenUser(user)
	return result
}

func (u UserUsecaseImp) GetLoginUser(username string, user *model.Login){
	u.repo.GetUser(username, user)
}