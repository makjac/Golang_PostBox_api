package main

import (
	"log"
	"net/http"

	database "github.com/PBB-api/config"
	"github.com/PBB-api/controller"
	mails "github.com/PBB-api/mail"
	"github.com/PBB-api/middlewares"
	"github.com/PBB-api/routers"
	"github.com/PBB-api/service"

	"github.com/gin-gonic/gin"
)

var (
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {
	//Connect to the postgres batabase
	database.SetupDB()

	mails.VerficationMail("maks.jackowski@wp.pl", "makjac", "wp.pl")

	//init router
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// Route Handlers; endpoints
	routers.RtrSetup(server)

	log.Fatal(server.Run(":8080"))
}
