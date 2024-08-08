package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.gtld.test/realTimeApp/app/models"
	"rest.gtld.test/realTimeApp/config"
	"rest.gtld.test/realTimeApp/database"

	// "github.com/olahol/melody"
	"github.com/gorilla/websocket"
)

type ginServer struct {
	app  *gin.Engine
	db   database.Database
	conf *config.Config
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
	  // Allow all connections
	  return true
	},
  }

func NewGinServer(conf *config.Config, db database.Database) Server {
	ginApp := gin.Default()

	return &ginServer{
		app:  ginApp,
		db:   db,
		conf: conf,
	}
}

func (s *ginServer) Start() {
	s.app.GET("v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "Health care",
		})
	})

	s.app.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
		  c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		  })
		  return
		}
		// Handle WebSocket connection
		go handleWebSocketConnection(conn)
	  })

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Run(serverUrl)
}


func handleWebSocketConnection(conn *websocket.Conn){
	for {
		// Read message from the client
		var weather model.Weather
		_, message, err := conn.ReadMessage()
		
		if err != nil {
		  conn.Close()
		  break
		}

		json.Unmarshal(message, &weather)
		fmt.Println(weather)
	  }
}