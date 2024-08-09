package repository

import (
	"rest.gtld.test/realTimeApp/app/entities"
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
