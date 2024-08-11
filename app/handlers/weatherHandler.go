package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"rest.gtld.test/realTimeApp/app/model"
	"rest.gtld.test/realTimeApp/app/usecases"
)

type weatherHandler struct {
	WeatherUsecaseImp *usecases.WeatherUsecaseImp
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

var mu sync.Mutex

func NewWeatherHandler(weatherUsecase *usecases.WeatherUsecaseImp) *weatherHandler {
	return &weatherHandler{
		WeatherUsecaseImp: weatherUsecase,
	}
}

func (w *weatherHandler) HandleWebSocketConnection(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	username, ok := c.Get("username")
	userstring := username.(string)
	if ok{
		mu.Lock()
		usecases.ActiveNode[userstring] = true
		mu.Unlock()
	}
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
	mu.Lock()
	usecases.ActiveNode[userstring] = false
	mu.Unlock()
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
