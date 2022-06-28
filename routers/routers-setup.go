package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RtrSetup(router *gin.Engine) {

	router.GET("/", welcome)
	router.NoRoute(notFound)
	PersonRouter(router)
	PostmachineRouter(router)
	jwtRouters(router)
	authRouters(router)

}

func welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome-page.html", gin.H{
		"staus":   http.StatusOK,
		"message": "oops... something is wrong",
	})
	return
}

func notFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"staus":   http.StatusNotFound,
		"message": "oops... something is wrong",
	})
	return
}
