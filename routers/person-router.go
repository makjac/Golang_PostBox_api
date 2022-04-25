package routers

import (
	"github.com/PBB-api/controller"
	"github.com/gin-gonic/gin"
)

func PersonRouter(router *gin.Engine) {
	router.GET("/persons", controller.GetAllPersons)
}
