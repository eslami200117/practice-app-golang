package usecases

import (
	"container/ring"
	"encoding/json"
	"fmt"
	"sync"

	"rest.gtld.test/realTimeApp/app/entities"
	"rest.gtld.test/realTimeApp/app/model"
	repository "rest.gtld.test/realTimeApp/app/repositories"
)

type WeatherUsecaseImp struct {
	weatherQueueRepo *ring.Ring
	weatherRepo      *repository.WeatherPostgresRepo
	mu               sync.Mutex
}

var (
	wr *ring.Ring = ring.New(1024)
	ActiveNode = make(map[string]bool)
)

func NewWeatherUseImp(weatherRepo *repository.WeatherPostgresRepo) *WeatherUsecaseImp {
	return &WeatherUsecaseImp{
		weatherQueueRepo: wr,
		weatherRepo:      weatherRepo,
		mu:               sync.Mutex{},
	}
}

func (wu *WeatherUsecaseImp) WeatherDataProcessing(w *model.Weather) error {

	insertWeatherData := &entities.WeatherEntity{
		Longitude:   w.Longitude,
		Latitude:    w.Latitude,
		WindSpeed:   w.WindSpeed,
		Temperature: w.Temperature,
		Rain:        w.Rain,
	}

	if err := wu.weatherRepo.InserWeatherData(insertWeatherData); err != nil {
		return err
	}

	wu.mu.Lock()
	wu.weatherQueueRepo.Value = *w
	wu.weatherQueueRepo = wu.weatherQueueRepo.Next()
	wu.mu.Unlock()
	return nil

}

func (wu *WeatherUsecaseImp) RainProccesin(lng, lat float64) float64 {
	i := 0.0
	re := 0.0
	wu.mu.Lock()
	wu.weatherQueueRepo.Do(func(p interface{}) {
		if p != nil {
			if v, ok := p.(model.Weather); ok {
				if v.Longitude > (lng-5) && v.Longitude < (lng+5) {
					if v.Latitude > (lat-5) && v.Latitude < (lat+5) { // Corrected to check Latitude
						re += v.Rain
						i++
					}
				}
			}
		}
	})
	wu.mu.Unlock()

	if i == 0 {
		return 0
	}

	return re / i
}

func (wu *WeatherUsecaseImp) LastValue() string {
	lastValue := wu.weatherQueueRepo.Prev().Value
	if lastValue == nil {
		return "empy"
	}
	jsondata, err := json.Marshal(lastValue)
	if err != nil {
		fmt.Println("error in marshal last value", err)
	}
	return string(jsondata)
	// return "empy2"
}
