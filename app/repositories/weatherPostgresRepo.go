package repository

import (
	"errors"
	"time"

	// "golang.org/x/exp/rand"
	"gorm.io/gorm"
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

func (pr *WeatherPostgresRepo) GetUser(username string, log *model.Login){
	var user model.User 
	pr.db.GetDb().Where("username=?", username).Find(&user)
	log.Password = user.Password
	log.Username = user.Username
}

func (pr *WeatherPostgresRepo) AuthenUser(in *model.Login) bool {
	var user model.User
	pr.db.GetDb().First(&user, "username= ?", in.Username)
	return user.Password == in.Password
}

func (pr *WeatherPostgresRepo) UpdateNodeStatus(username string, status bool){
	var node entities.Nodes
	pr.db.GetDb().First(&node, "username= ?", username)
	node.Status = status
	node.LastUpdata = time.Now()
	pr.db.GetDb().Save(&node)
}

func (pr *WeatherPostgresRepo) UpdateLastLogin(username string, lastLoginTime time.Time) {
	var user entities.User
	pr.db.GetDb().First(&user, "username= ?", username)
	user.LastLogin = lastLoginTime
	pr.db.GetDb().Save(&user)
}

func (pr *WeatherPostgresRepo) GetAllUser() []entities.User {
	var allUser []entities.User

	pr.db.GetDb().Find(&allUser)
	return allUser
}

func (pr *WeatherPostgresRepo) GetAllNode() []entities.Nodes {
	var allNode []entities.Nodes

	pr.db.GetDb().Find(&allNode)
	return allNode
}

func (pr *WeatherPostgresRepo) IsAdmin(username string) bool {
	var user entities.User
	pr.db.GetDb().First(&user, "username= ?", username)
	return user.Role == "superviser" 
}


func (pr *WeatherPostgresRepo) AddUser(username string, password string) error {
	var user entities.User

	if err := pr.db.GetDb().Where("username = ?", username).First(&user).Error; err == nil {
		return errors.New("user already exists")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	err := pr.db.GetDb().Create(&entities.User{
		Username:  username,
		Password:  password,
		Role:      "employee",
		LastLogin: time.Now(),
	}).Error

	return err
}

func (pr *WeatherPostgresRepo) AddSource(addSource string, password string) error {
	var node entities.Nodes

	if err := pr.db.GetDb().Where("username = ?", addSource).First(&node).Error; err == nil {
		return errors.New("source already exists")
	} else if err != gorm.ErrRecordNotFound {
		return err 
	}

	err := pr.db.GetDb().Create(&entities.Nodes{
		Username:  addSource,
		Password:  password,
		Role:      "worker",
		Status:    false,
		LastUpdata: time.Now(),
	}).Error

	return err
}
