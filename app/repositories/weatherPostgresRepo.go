package repository

import (
	"rest.gtld.test/realTimeApp/app/entities"
	"rest.gtld.test/realTimeApp/app/model"
	"rest.gtld.test/realTimeApp/database"
)

type WeatherPostgresRepo struct {
	db database.Database
}

func NewWeatherPostgresRepo(db database.Database) *WeatherPostgresRepo {
	return &WeatherPostgresRepo{
		db: db,
	}
}

func (pr *WeatherPostgresRepo) InserWeatherData(data *entities.WeatherEntity) error {
	result := pr.db.GetDb().Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *WeatherPostgresRepo) AuthenNode(in *model.Login) bool{
	var node model.Node
	pr.db.GetDb().First(&node, "username= ?", in.Username)
	return node.Password == in.Password
}

func (pr *WeatherPostgresRepo) GetNode(username string, user *model.Login){
	var node model.Node 
	pr.db.GetDb().Where("username=?", username).Find(&node)
	user.Password = node.Password
	user.Username = node.Username
}

func (pr * WeatherPostgresRepo) AuthenUser(in *model.Login) bool {
	var user model.User
	pr.db.GetDb().First(&user, "username= ?", in.Username)
	return user.Password == in.Password
}