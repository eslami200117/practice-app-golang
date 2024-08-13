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

var (
	mu sync.Mutex
	selectedNode string
)

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
	w.WeatherUsecaseImp.UpdateNodeStatus(userstring, true)

	for {
		var weather model.Weather
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			break
		}
		json.Unmarshal(message, &weather)
		if userstring == selectedNode {
			usecases.SelectedValue = weather
		}
		w.WeatherUsecaseImp.WeatherDataProcessing(&weather)
	}

	w.WeatherUsecaseImp.UpdateNodeStatus(userstring, false)
	mu.Lock()
	usecases.ActiveNode[userstring] = false
	mu.Unlock()
}

func (w *weatherHandler) HaddleUserRecPrc(c *gin.Context) {
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

func (u weatherHandler) ListHandler(c *gin.Context){
	// need modification
	c.JSON(http.StatusOK, gin.H{
		"nodes":	usecases.ActiveNode,
	})
}

func (u weatherHandler) GetNodeHandler(c *gin.Context){
	selectedNode = c.Query("node")

	activitie, ok := usecases.ActiveNode[selectedNode]
	if ok && activitie {
		c.JSON(http.StatusOK, gin.H{
			"data":	usecases.SelectedValue,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error":	"unactive node",
		})
	}
}