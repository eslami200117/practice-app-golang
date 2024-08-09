package usecases

import (
	"rest.gtld.test/realTimeApp/app/model"
)

type WeatherUsecase interface {
	WeatherDataProcessing(*model.Weather)
}
