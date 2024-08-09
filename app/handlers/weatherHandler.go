package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"rest.gtld.test/realTimeApp/app/model"
	"rest.gtld.test/realTimeApp/app/usecases"
)

type weatherHandler struct {
	WeatherUsecaseImp *usecases.WeatherUsecaseImp
}

func NewWeatherHandler(weatherUsecase *usecases.WeatherUsecaseImp) *weatherHandler {
	return &weatherHandler{
		WeatherUsecaseImp: weatherUsecase,
	}
}

func (w *weatherHandler) HandleWebSocketConnection(conn *websocket.Conn) {
	for {
		var weather model.Weather
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			break
		}
		json.Unmarshal(message, &weather)
		w.WeatherUsecaseImp.WeatherDataProcessing(&weather)
	}
}

func (w *weatherHandler) HnadleUserRecPrc(c *gin.Context) {
	lngStr := c.Query("lng")
    latStr := c.Query("lat")

    lng, err := strconv.ParseFloat(lngStr, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
        return
    }

    lat, err := strconv.ParseFloat(latStr, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
        return
    }

    result := w.WeatherUsecaseImp.RainProccesin(lng, lat)

	c.JSON(http.StatusOK, gin.H{
		"avr rain":   result,
		"last value": w.WeatherUsecaseImp.LastValue(),
	})

}
