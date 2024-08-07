package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.gtld.test/realTimeApp/config"
	"rest.gtld.test/realTimeApp/database"
)

type ginServer struct {
	app  *gin.Engine
	db   database.Database
	conf *config.Config
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

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Run(serverUrl)
}
