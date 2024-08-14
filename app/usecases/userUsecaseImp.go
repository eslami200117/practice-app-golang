package usecases

import (
	"time"

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

func (u *UserUsecaseImp) AuthenticateUser(c *gin.Context, user *model.Login) bool {
	result := u.repo.AuthenUser(user)
	return result
}

func (u *UserUsecaseImp) GetLoginUser(username string, user *model.Login){
	u.repo.GetUser(username, user)
}

func (u *UserUsecaseImp) UpdateLastLogin(username string, lastLoginTime time.Time){
	u.repo.UpdateLastLogin(username, lastLoginTime)
}

func (u *UserUsecaseImp) GetAllUser() []model.APIuser{
	allUser := u.repo.GetAllUser()
	var allAPIuser []model.APIuser
	for _, u := range allUser{
		value := model.APIuser{
			Username: u.Username,
			LastLogin: timeAgo(u.LastLogin),
		}

		allAPIuser = append(allAPIuser, value)
	}
	return allAPIuser
}

func (u *UserUsecaseImp) IsAdmin(username string) bool {
	return u.repo.IsAdmin(username)
}

func (u *UserUsecaseImp) AddUser (username string, password string) error{
	return u.repo.AddUser(username, password)
}