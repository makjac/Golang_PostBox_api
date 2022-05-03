package main

import (
	"log"

	database "github.com/PBB-api/config"
	"github.com/PBB-api/middlewares"
	"github.com/PBB-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//Connect to the postgres batabase
	database.SetupDB()

	//init router
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	// Route Handlers; endpoints
	routers.RtrSetup(server)

	log.Fatal(server.Run(":8080"))
}
