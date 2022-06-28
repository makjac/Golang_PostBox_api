package main

import (
	"log"
	"time"

	database "github.com/PBB-api/config"
	"github.com/PBB-api/cors"
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
	server.LoadHTMLGlob("templates/*.html")

	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	routers.RtrSetup(server)

	log.Fatal(server.Run(":8080"))
}
