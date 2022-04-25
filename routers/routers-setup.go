package routers

import (
	"net/http"

	"github.com/PBB-api/controller"
	"github.com/PBB-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RtrSetup(router *gin.Engine) {

	router.GET("/", welcome)
	router.NoRoute(notFound)
	PersonRouter(router)

	apiRouters := router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRouters.GET("/history/:id", controller.GetParcelHistory)
		apiRouters.GET("/adress/:id", controller.GetAdressID)
	}
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"staus":   202,
		"message": "Welcome",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"staus":   404,
		"message": "oops... something is wrong",
	})
	return
}
