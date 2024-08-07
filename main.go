package main

import (
	"rest.gtld.test/realTimeApp/config"
	"rest.gtld.test/realTimeApp/database"
	"rest.gtld.test/realTimeApp/server"
)

func main(){
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	router := server.NewGinServer(conf, db)
	router.Start()
}
