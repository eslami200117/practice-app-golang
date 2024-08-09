package usecases

import (
	"sync"

	"rest.gtld.test/realTimeApp/app/model"
)

type WeatherUsecaseImp struct {
	weatherRepo []model.Weather
	mu sync.Mutex
}


var wr []model.Weather

func NewWeatherUseImp()(*WeatherUsecaseImp){
	return &WeatherUsecaseImp{
		weatherRepo: wr,
		mu: sync.Mutex{},
	}
}

func (wu *WeatherUsecaseImp) WeatherDataProcessing(w *model.Weather){
	wu.mu.Lock()
	wu.weatherRepo = append(wu.weatherRepo, *w)
	wu.mu.Unlock()
}

func (wu WeatherUsecaseImp) RainProccesin(lng, lat float64)(float64){
	i := 0.0
	re := 0.0
	for _, v := range wu.weatherRepo[(len(wu.weatherRepo)-100):]{
		if v.Longitude > (lng - 5) &&  v.Longitude < (lng + 5) {
			if v.Latitude > (lng - 5) &&  v.Latitude < (lng + 5){
				re += v.Rain
				i++
			}
		}
	}
	re/=i

	return re
}
