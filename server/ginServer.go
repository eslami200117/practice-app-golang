package server

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
	handler "rest.gtld.test/realTimeApp/app/handlers"
	repository "rest.gtld.test/realTimeApp/app/repositories"
	"rest.gtld.test/realTimeApp/app/usecases"

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
	s.initialWeatherHandler()
	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Run(serverUrl)
}

func (s *ginServer) initialWeatherHandler() {
	repo := repository.NewWeatherPostgresRepo(s.db)
	weatherUsecase := usecases.NewWeatherUseImp(repo)
	weatherHandler := handler.NewWeatherHandler(weatherUsecase)

	nodeUsercase := usecases.NewNodeImp(repo)
	nodeHandler := handler.NewNodeHandler(nodeUsercase)

	s.app.GET("/ws", func(c *gin.Context) {
		weatherHandler.HandleWebSocketConnection(c)
	})

	s.app.POST("/login", func(c *gin.Context){
		nodeHandler.HnadleLogin(c)
	})

	s.app.GET("test", func(c *gin.Context) {
		weatherHandler.HnadleUserRecPrc(c)
	})
}
