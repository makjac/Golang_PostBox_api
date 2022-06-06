package main

import (
	"log"

	database "github.com/PBB-api/config"
	//mails "github.com/PBB-api/mail"
	"github.com/PBB-api/middlewares"
	"github.com/PBB-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//Connect to the postgres batabase
	database.SetupDB()

	//śmails.VerficationMail("veslymass@gmail.com", "maks", "http://makjac.pl:8080/activated")

	//init router
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	// Route Handlers; endpoints
	server.LoadHTMLGlob("templates/*.html")
	routers.RtrSetup(server)

	log.Fatal(server.Run(":8080"))
}
