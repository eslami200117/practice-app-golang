package main

import (
	"rest.gtld.test/realTimeApp/app/entities"
	"rest.gtld.test/realTimeApp/config"
	"rest.gtld.test/realTimeApp/database"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	weatherMigrate(db)
	nodeMigrate(db)
}

func weatherMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.WeatherEntity{})
	db.GetDb().Create(
		&entities.WeatherEntity{
			Longitude:   0,
			Latitude:    0,
			WindSpeed:   0,
			Temperature: 0,
			Rain:        0,
		},
	)
}

func nodeMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Nodes{})
	db.GetDb().Create(
		&entities.Nodes{
			Id: 1001,
			Username: "test",
			Password: "password",
			Role: "worker",
		},
	)


}