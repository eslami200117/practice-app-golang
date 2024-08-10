package repository

import "rest.gtld.test/realTimeApp/app/entities"

type WeatherRepo interface {
	InserWeatherData(data entities.WeatherEntity) error
	Authen(username string, password string) bool
}
